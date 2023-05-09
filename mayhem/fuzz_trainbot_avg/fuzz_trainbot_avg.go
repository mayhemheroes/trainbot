package fuzz_trainbot_avg

import (
    "bytes"
    "image"

    "github.com/jo-m/trainbot/pkg/avg"
    "github.com/jo-m/trainbot/pkg/imutil"
)

func mayhemit(data []byte) int {

    if len(data) > 2 {
        num := int(data[0])
        data = data[1:]
        
        switch num {
            
            case 0:
                img, _, err := image.Decode(bytes.NewReader(data))
                if err != nil {
                    return 0
                }
                result := imutil.ToGray(img)

                avg.Gray(result)
                return 0

            case 1:
                img, _, err := image.Decode(bytes.NewReader(data))
                if err != nil {
                    return 0
                }
                result := imutil.ToGray(img)

                avg.GraySlow(result)
                return 0

            case 2:
                img, _, err := image.Decode(bytes.NewReader(data))
                if err != nil {
                    return 0
                }
                result := imutil.ToRGBA(img)

                avg.RGBA(result)
                return 0

            case 3:
                img, _, err := image.Decode(bytes.NewReader(data))
                if err != nil {
                    return 0
                }
                result := imutil.ToRGBA(img)

                avg.RGBASlow(result)
                return 0
        }
    }
    return 0
}

func Fuzz(data []byte) int {
    _ = mayhemit(data)
    return 0
}