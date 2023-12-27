// inspired by https://github.com/chrisveness/latlon-geohash/blob/master/latlon-geohash.js

package geohash

import (
	"math"
	"strings"
)

const (
    base32 = "0123456789bcdefghjkmnpqrstuvwxyz"
)

type Box struct {
    Min, Max float64
}

func Encode(lat, lon float64, precision int) string {
    var geohash []rune
    var bits int = 0
    var bit uint = 0
    even := true
    latBox := Box{-90.0, 90.0}
    lonBox := Box{-180.0, 180.0}

    for len(geohash) < precision {
        if even {
            mid := (lonBox.Min + lonBox.Max) / 2
            if lon > mid {
                bit |= 1 << uint(bits)
                lonBox.Min = mid
            } else {
                lonBox.Max = mid
            }
        } else {
            mid := (latBox.Min + latBox.Max) / 2
            if lat > mid {
                bit |= 1 << uint(bits)
                latBox.Min = mid
            } else {
                latBox.Max = mid
            }
        }

        even = !even
        if bits < 4 {
            bits++
        } else {
            geohash = append(geohash, rune(base32[bit]))
            bits = 0
            bit = 0
        }
    }

    return string(geohash)
}

func Decode(geohash string) (lat, lon float64) {
    even := true
    latBox := Box{-90.0, 90.0}
    lonBox := Box{-180.0, 180.0}

    for _, c := range geohash {
        i := strings.Index(base32, string(c))
        for mask := 16; mask != 0; mask >>= 1 {
            if even {
                if i&mask != 0 {
                    lonBox.Min = (lonBox.Min + lonBox.Max) / 2
                } else {
                    lonBox.Max = (lonBox.Min + lonBox.Max) / 2
                }
            } else {
                if i&mask != 0 {
                    latBox.Min = (latBox.Min + latBox.Max) / 2
                } else {
                    latBox.Max = (latBox.Min + latBox.Max) / 2
                }
            }
            even = !even
        }
    }

    lat = (latBox.Min + latBox.Max) / 2
    lon = (lonBox.Min + lonBox.Max) / 2

    return lat, lon
}

func Distance(lat1, lon1, lat2, lon2 float64) float64 {
    rad := math.Pi / 180
    lat1 *= rad
    lon1 *= rad
    lat2 *= rad
    lon2 *= rad

    a := math.Sin((lat2-lat1)/2)*math.Sin((lat2-lat1)/2) +
        math.Cos(lat1)*math.Cos(lat2)*math.Sin((lon2-lon1)/2)*math.Sin((lon2-lon1)/2)
    c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

    return 6371e3 * c // Earth radius in meters
}