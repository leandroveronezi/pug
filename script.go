package pug

type TScript struct {
	TDomElement
}

func (_this TScript) New(Id string) *TScript {

	_this.TDomElement.id = Id
	_this.TDomElement.elementType = TDomElemenScript
	_this.TDomElement.Create()

	return &_this
}

func (_this TScript) Id() string {
	return _this.id
}

func (_this *TScript) Src(src string) *TScript {
	_this.TDomElement.el.Set("src", src)
	return _this
}
