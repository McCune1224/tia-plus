package main

import (
	"fmt"
	"strconv"
	"strings"
)

type IPAddress string

// Split the 4 octets of the IP address into a list of ints
func (i *IPAddress) Split() ([]int, error) {
	chunks := strings.Split(string(*i), ".")
	ipList := []int{}
	if len(chunks) != 4 {
		return ipList, fmt.Errorf("Invalid. Expected 4 '.' but got %d", len(chunks))
	}
	for i, octet := range chunks {
		intIP, err := strconv.Atoi(octet)
		if err != nil {
			return ipList, err
		}
		if intIP > 255 || intIP < 0 {
			return ipList, fmt.Errorf("Octect #%d has an invalid decimal value of %d in %s", i, intIP, octet)
		}
		ipList = append(ipList, intIP)
	}
	return ipList, nil
}

// Split the 4 octets of the IP address into a list of strings
func (i *IPAddress) SplitS() ([]string, error) {
	chunks := strings.Split(string(*i), ".")
	ipList := []string{}
	for i, octet := range chunks {
		intIP, err := strconv.Atoi(octet)
		if err != nil {
			return ipList, err
		}
		if intIP > 255 || intIP < 0 {
			return ipList, fmt.Errorf("Octect #%d has an invalid decimal value of %d in %s", i, intIP, octet)
		}
		ipList = append(ipList, octet)
	}
	return ipList, nil
}

// Follow the "Magic Number" method https://www.youtube.com/watch?v=P1ROXMLjL04
func CalculateSubnetID(address IPAddress, subnetMask IPAddress) (IPAddress, error) {
	result := IPAddress("")
	rawIPString := []string{}
	ipBlocks, err := address.Split()
	subnetBlocks, err := subnetMask.Split()
	if err != nil {
		return result, err
	}
	for i := 0; i < len(ipBlocks); i++ {
		ipOctet := ipBlocks[i]
		maskOctet := subnetBlocks[i]
		switch maskOctet {
		case 255:
			rawIPString = append(rawIPString, strconv.Itoa(ipOctet))
		case 0:
			rawIPString = append(rawIPString, "0")
		default:
			magicNumber := 256 - maskOctet
			subnetIDOctet := (ipOctet / magicNumber) * magicNumber
			rawIPString = append(rawIPString, strconv.Itoa(subnetIDOctet))
		}
	}
	result = IPAddress(strings.Join(rawIPString, "."))
	return result, nil
}

func CalculateBroadcastAddress(subnetID IPAddress, subnetMask IPAddress) (IPAddress, error) {
	result := IPAddress("")
	rawIPString := []string{}
	subnetIDBlock, err := subnetID.Split()
	subnetMaskBlocks, err := subnetMask.Split()

	if err != nil {
		return result, err
	}

	for i := 0; i < len(subnetIDBlock); i++ {
		subnetIDOctet := subnetIDBlock[i]
		maskOctet := subnetMaskBlocks[i]
		switch maskOctet {
		case 255:
			rawIPString = append(rawIPString, strconv.Itoa(subnetIDOctet))
		case 0:
			rawIPString = append(rawIPString, "255")
		default:
			magicNumber := 256 - maskOctet
			braodcastOctet := subnetIDOctet + magicNumber - 1
			fmt.Println(magicNumber, subnetIDOctet, braodcastOctet)
			rawIPString = append(rawIPString, strconv.Itoa(braodcastOctet))
		}
	}

	result = IPAddress(strings.Join(rawIPString, "."))
	return result, nil
}

func main() {
	target := IPAddress("165.245.77.14")
	subnetMask := IPAddress("255.255.240.0")

	subnetID, err := CalculateSubnetID(target, subnetMask)
	if err != nil {
		fmt.Println(err)
		return
	}
	broadcast, err := CalculateBroadcastAddress(subnetID, subnetMask)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("TARGET IP: %s\nTARGET SUBNET:%s\nSUBNET ID:%s\nBroadcast Address:%s\n", target, subnetMask, subnetID, broadcast)
}
