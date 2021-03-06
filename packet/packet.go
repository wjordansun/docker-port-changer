package packet

import (
	"fmt"
	"log"
	"os"
	"portchanger/badgerstuff"
	"portchanger/docker"
	"portchanger/ipfs"
	"portchanger/ports"
	"strings"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"github.com/google/gopacket/pcapgo"
)

var (
    device       		string = "docker0"
    snapshot_len 		int32  = 1024
		snapshotLen  		uint32 = 1024
    promiscuous  		bool   = false
    err          		error
    timeout      		time.Duration = 1 * time.Second
    handle       		*pcap.Handle
		packetCount	 	 	int = 0
		packetsPerFile	int = 100
		pcapFile				string  = "sample.pcap"
		CID 						string = ""
		honeypot3				bool = false
)

func writeFile(packet gopacket.Packet, w *pcapgo.Writer) {

	w.WritePacket(packet.Metadata().CaptureInfo, packet.Data())
	packetCount++

	if packetCount > packetsPerFile {
		CID = ipfs.Add(pcapFile)
		os.Exit(0)
	}
}

func OpenFile() {

	handle, err = pcap.OpenOffline(pcapFile)
  if err != nil { log.Fatal(err) }
  defer handle.Close()

  // Loop through packets in file
  packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
  for packet := range packetSource.Packets() {
    fmt.Println(packet)
  }
}

func Listen() {
    // Open device
    handle, err = pcap.OpenLive(device, snapshot_len, promiscuous, timeout)
    if err != nil {
        log.Fatal(err)
    }
    defer handle.Close()

		//Open output pcap file and write header 
		// f, _ := os.Create("test.pcap")
		// w := pcapgo.NewWriter(f)
		// w.WriteFileHeader(snapshotLen, layers.LinkTypeEthernet)
		// defer f.Close()

    // Set filter
    // var filter string = "tcp and port 8080"
    // err = handle.SetBPFFilter(filter)
    // if err != nil {
    //     log.Fatal(err)
    // }
    // fmt.Println("Only capturing TCP port 8080 packets.")
		
		//reader := bufio.NewReader(os.Stdin)
		
    packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

		//Loop:
    	for packet := range packetSource.Packets() {
				// cmdString, err := reader.ReadString('\n')
				// commandStr := strings.TrimSuffix(cmdString, "\n")
				// badgerstuff.Handle(err)
				// if commandStr == "q" {
				// 	break Loop
				// }

				pac := packet.String()
				//fmt.Println(packet)
        fmt.Println(strings.Contains(pac, "RST=true"))
				//writeFile(packet, w)
				if strings.Contains(pac, "RST=true") {
					packetCount = 101
					fmt.Println(packet)
					switch badgerstuff.ProductionNum() {
					case 1:
						if honeypot3 {
							docker.Stop("honeypot3")
							docker.Stop("production1")
         	 		docker.Start("honeypot1")
							docker.Start("production2")
						} else {
							time.Sleep(2 * time.Second)
							docker.Stop("mongodb")
							docker.Start("mongodb1")
							docker.Start("mongodbh")

							time.Sleep(2 * time.Second)

							docker.Stop("production1")
         	 		docker.Start("honeypot1")
							docker.Start("production2")
							badgerstuff.SetProductionNum(2)
							
							ports.Display()

							time.Sleep(2 * time.Second)
							os.Exit(0)
						}
					case 2:
						docker.Stop("production2")
						docker.Stop("honeypot1")
          	docker.Start("honeypot2")
						docker.Start("production3")
						badgerstuff.SetProductionNum(3)
					case 3:
						docker.Stop("production3")
						docker.Stop("honeypot2")
          	docker.Start("honeypot3")
						docker.Start("production1")
						honeypot3 = true
						badgerstuff.SetProductionNum(1)
					}
					

				}
    }

}