package controllers

import "github.com/robfig/revel"
import "github.com/gokyle/hotp"
import "encoding/base64"
import "errors"

const numDigits = 6

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) storeHOTP(otp *hotp.HOTP) error {
	out, err := hotp.Marshal(otp)
	if err != nil {
		c.Flash.Error("Oh no! I couldn't store the HOTP key value!")
		revel.ERROR.Printf("failed to store HOTP key value: %v", err)
		c.FlashParams()
		return err
	}

	c.Session["otp"] = base64.StdEncoding.EncodeToString(out)
	return nil
}

func (c App) loadHOTP() (*hotp.HOTP, error) {
	encoded, ok := c.Session["otp"]
	if !ok {
		c.Flash.Error("Oh no! I couldn't store the HOTP key value!")
		revel.ERROR.Println("HOTP key value not present")
		c.FlashParams()
		return nil, errors.New("failed to restore HOTP")
	}

	in, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		c.Flash.Error("Oh no! I couldn't store the HOTP key value!")
		revel.ERROR.Printf("failed to store HOTP key value: %v", err)
		c.FlashParams()
		return nil, err
	}

	otp, err := hotp.Unmarshal(in)
	if err != nil {
		c.Flash.Error("Oh no! I couldn't store the HOTP key value!")
		revel.ERROR.Printf("failed to store HOTP key value: %v", err)
		c.FlashParams()
		return nil, err
	}
	return otp, err
}

func (c App) NewHotp(name string) revel.Result {
	otp, err := hotp.GenerateHOTP(numDigits, false)
	if err != nil {
		c.Flash.Error("Sorry - I couldn't generate a new HOTP!")
		c.FlashParams()
		return c.Redirect(App.Index)
	}

	qr, err := otp.QR(name)
	if err != nil {
		c.Flash.Error("Sorry - I couldn't generate a new HOTP!")
		c.FlashParams()
		return c.Redirect(App.Index)
	}

	png := base64.StdEncoding.EncodeToString(qr)
	code0 := otp.OTP()
	code1 := otp.OTP()

	err = c.storeHOTP(otp)
	if err != nil {
		return c.Redirect(App.Index)
	}
	c.Session["name"] = name
	return c.Render(name, png, code0, code1)
}

func (c App) EnterCode() revel.Result {
	name, ok := c.Session["name"]
	if !ok || c.Session["otp"] == "" {
		c.Flash.Error("Woops! You need to start from the beginning.")
		c.FlashParams()
		return c.Redirect(App.Index)
	}

	return c.Render(name)
}

func (c App) CheckCode(code string) revel.Result {
	_, ok := c.Session["name"]
	if !ok || c.Session["otp"] == "" {
		c.Flash.Error("Woops! You need to start from the beginning.")
		c.FlashParams()
		return c.Redirect(App.Index)
	} else if code == "" {
		c.Flash.Error("Nice try, but you have to enter a code for this to work.")
		c.FlashParams()
		return c.Redirect(App.EnterCode)
	}

	otp, err := c.loadHOTP()
	if err != nil {
		return c.Redirect(App.Index)
	}

	if otp.Scan(code, 3) {
		c.Flash.Success("Great! That code worked.")
	} else {
		c.Flash.Error("Looks like that was an invalid code.")
	}
	err = c.storeHOTP(otp)
	if err != nil {
		return c.Redirect(App.Index)
	}
	c.FlashParams()
	return c.Redirect(App.EnterCode)
}

func (c App) IntegrityCheck() revel.Result {
	name, ok := c.Session["name"]
	if !ok || c.Session["otp"] == "" {
		c.Flash.Error("Woops! You need to start from the beginning.")
		c.FlashParams()
		return c.Redirect(App.Index)
	}
	otp, err := c.loadHOTP()
	if err != nil {
		return c.Redirect(App.Index)
	}

	code, counter := otp.IntegrityCheck()
	counter--
	return c.Render(name, code, counter)
}
