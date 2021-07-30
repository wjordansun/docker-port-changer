package packet

import (
	"fmt"
	"log"
	"os"
	"portchanger/docker"
	"portchanger/ipfs"
	"strings"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
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

		// Open output pcap file and write header 
		f, _ := os.Create("test.pcap")
		w := pcapgo.NewWriter(f)
		w.WriteFileHeader(snapshotLen, layers.LinkTypeEthernet)
		defer f.Close()

    // Set filter
    // var filter string = "tcp and port 8080"
    // err = handle.SetBPFFilter(filter)
    // if err != nil {
    //     log.Fatal(err)
    // }
    // fmt.Println("Only capturing TCP port 8080 packets.")

    packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
    for packet := range packetSource.Packets() {
				pac := packet.String()
				//fmt.Println(packet)
        //fmt.Println(strings.Contains(pac, "RST=true"))
				writeFile(packet, w)
				if strings.Contains(pac, "RST=true") {

					fmt.Println(packet)
					
					docker.Stop()

          docker.Start()
        }
    }

}