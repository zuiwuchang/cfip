package cf

import (
	"crypto/rand"
	mathRand "math/rand"
	"net"
)

type IPRange struct {
	ip    []net.IP
	ipnet []*net.IPNet
}

func (r *IPRange) AddCIDR(s string) (e error) {
	ip, ipnet, e := net.ParseCIDR(s)
	if e != nil {
		return
	}
	ones, bits := ipnet.Mask.Size()
	if ones == bits {
		return
	}
	r.ip = append(r.ip, ip)
	r.ipnet = append(r.ipnet, ipnet)
	return
}

// 返回一個隨機 ip
func (r IPRange) Random() net.IP {
	size := len(r.ipnet)
	if size == 0 {
		return nil
	}

	for {
		i := mathRand.Intn(size)
		ip := r.ip[i]
		ipnet := r.ipnet[i]

		ones, bits := ipnet.Mask.Size()
		if bits == 32 {
			ip = ip.To4()
		}

		b := make([]byte, bits/8)
		rand.Read(b)

		i = ones / 8
		copy(b, ip[:i])
		switch ones % 8 {
		case 1:
			b[i] &= 0x7f
			b[i] |= ip[i]
		case 2:
			b[i] &= 0x3f
			b[i] |= ip[i]
		case 3:
			b[i] &= 0x1f
			b[i] |= ip[i]
		case 4:
			b[i] &= 0xf
			b[i] |= ip[i]
		case 5:
			b[i] &= 0x7
			b[i] |= ip[i]
		case 6:
			b[i] &= 0x3
			b[i] |= ip[i]
		case 7:
			b[i] &= 0x1
			b[i] |= ip[i]
		}
		if ones%8 != 0 {
			b[i] &= ip[i]
		}

		if !ipnet.Contains(net.IP(b)) {
			continue
			// log.Fatalln(net.IP(b))
		}
		return net.IP(b)
	}
}
