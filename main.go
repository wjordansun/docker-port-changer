package main

import (
	"fmt"
	"log"
	"portchanger/docker"
	"strings"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

var (
    device       string = "docker0"
    snapshot_len int32  = 1024
    promiscuous  bool   = false
    err          error
    timeout      time.Duration = 1 * time.Second
    handle       *pcap.Handle
)

func main() {
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
        fmt.Println(strings.Contains(pac, "RST=true"))
				if strings.Contains(pac, "RST=true") {

						go docker.Stop()
						time.Sleep(1 * time.Second)

            docker.Start()
        }
    }

}