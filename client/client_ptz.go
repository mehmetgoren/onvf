package client

import (
	"github.com/use-go/onvif/ptz"
	onvif2 "github.com/use-go/onvif/xsd/onvif"
	"io/ioutil"
	"log"
)

func (o *OnvifClient) Ptz(x float64, y float64, zoom float64) (bool, error) {
	profiles, err := o.GetProfiles()
	if err != nil {
		return false, err
	}
	if len(profiles) == 0 {
		return false, nil
	}
	profileToken := profiles[0].Token

	profile := onvif2.ReferenceToken(profileToken)
	res, err := o.d.CallMethod(ptz.AbsoluteMove{
		ProfileToken: profile,
		Position: onvif2.PTZVector{
			PanTilt: onvif2.Vector2D{
				X:     x,
				Y:     y,
				Space: "http://www.onvif.org/ver10/tptz/PanTiltSpaces/TranslationGenericSpace",
			},
			Zoom: onvif2.Vector1D{
				X:     zoom,
				Space: "http://www.onvif.org/ver10/tptz/ZoomSpaces/TranslationGenericSpace",
			},
		},
	})
	log.Print(err)
	bs, _ := ioutil.ReadAll(res.Body)
	log.Printf("output %+v %s", res.StatusCode, bs)

	return true, nil
}
