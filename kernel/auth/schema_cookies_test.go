package auth

import (
	"fmt"
	auth2 "github.com/YuzuWiki/Pixivlee/kernel/auth"
	"testing"
)

func Test_TAuthCookie(t *testing.T) {
	sessId := "97684962_t4eNLFv7KgRJMoaeoTuwJF9qnIFd0sGU"
	cookies := fmt.Sprint("first_visit_datetime_pc=2024-09-23%2011%3A52%3A08; cc1=2024-09-23%2011%3A52%3A08; p_ab_id=2; p_ab_id_2=3; p_ab_d_id=1040258327; yuid_b=IVVXZgY; __cf_bm=ALeALZh9zmCVbMCdUmVmy57zpvB4_Wvdu58ZVkgW5GY-1727059928-1.0.1.1-DXzrjJQ.MzFsFAeoSgyP0KcQ_BXeeAU1TsdqbNjUHPJyZ4U6.cmJRHM17JM0B.cunUwvM_7fRH5ZLubwOmyfWIQ1u8FfLl7bJSKReSjW9YA; __utma=235335808.13391479.1727059926.1727059926.1727059926.1; __utmc=235335808; __utmz=235335808.1727059926.1.1.utmcsr=(direct)|utmccn=(direct)|utmcmd=(none); __utmv=235335808.|2=login%20ever=no=1^9=p_ab_id=2=1^10=p_ab_id_2=3=1^11=lang=zh=1; __utmt=1; __utmb=235335808.1.10.1727059926; cf_clearance=mJ0gyNHFAlFCuvXwTUS773RbxhvbCqb4Rlqu8wkzJ58-1727059929-1.2.1.1-RaZ4NrXhYyLigRWAKqPGpt0h6LDPSgUo8m_kSCQpA7uibHXYUM1tP0P_LQ1DGgAjTaTYgL7hLM8lOYz52ikK88MKUSkT3Zxxv5Rl2ZSamh5.PnmZ468j8IKRgxfhB3sdNUucoykSR7OZBwjpo0W55peXHJphjGYZ4dayvrEZavuGEXbHbdMtzOl3BkZWlcJF; _ga_75BBYNYN9J=GS1.1.1727059926.1.0.1727059931.0.0.0; _gid=GA1.2.2091318232; _gat_UA-1830249-3=1; _ga=GA1.1.862297305; PHPSESSID=" + sessId + "; device_token=20201b98bc42c54d319bab4e9f0452be; c_type=32; privacy_policy_agreement=0; privacy_policy_notification=0; a_type=0; b_type=0; _ga_MZ1NL4PHH0=GS1.1.1727059932.1.1.1727059965.0.0.0")

	a, err := auth2.NewAuthCookie(cookies).Do()
	if err != nil {
		t.Error(err)
		return
	}

	if a.SessId() != sessId {
		t.Errorf("parse error (not equal)")
		return
	}
	return
}
