package packet

import (
	"fmt"
	"log"
	"os"
	"portchanger/docker"
	"strings"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"github.com/google/gopacket/pcapgo"
)

var (
    device       string = "docker0"
    snapshot_len int32  = 1024
		snapshotLen  uint32 = 1024
    promiscuous  bool   = false
    err          error
    timeout      time.Duration = 1 * time.Second
    handle       *pcap.Handle
		packetCount	 int = 0
)

func writeFile(packet gopacket.Packet) {
	// Open output pcap file and write header 
	f, _ := os.Create("test.pcap")
	w := pcapgo.NewWriter(f)
	w.WriteFileHeader(snapshotLen, layers.LinkTypeEthernet)
	defer f.Close()

	w.WritePacket(packet.Metadata().CaptureInfo, packet.Data())
	packetCount++

	if packetCount > 100 {
		
	}
}

func Listen() {
    // Open device
    handle, err = pcap.OpenLive(device, snapshot_len, promiscuous, timeout)
    if err != nil {
        log.Fatal(err)
    }
    defer handle.Close()

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
				if strings.Contains(pac, "RST=true") {

					fmt.Println(packet)

					writeFile(packet)
					
					docker.Stop()

          docker.Start()
        }
    }

}