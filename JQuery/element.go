package JQuery

import "syscall/js"

func GetElement(Selector ...interface{}) *JQueryElement {

	element := js.Global().Call("$", Selector...)

	if element.Type() == js.TypeUndefined || element.Type() == js.TypeNull {
		return nil
	}

	return &JQueryElement{Element: element}

}

type JQueryElement struct {
	Element js.Value
}

/*
click()
Attaches/Triggers the click event
*/
func (this JQueryElement) Click(Value ...interface{}) js.Value {
	return this.Element.Call("click", Value...)
}

/*
dblclick()
Attaches/Triggers the double click event
*/
func (this JQueryElement) Dblclick(Value ...interface{}) js.Value {
	return this.Element.Call("dblclick", Value...)
}

/*
blur()
Attaches/Triggers the blur event
*/
func (this JQueryElement) Blur(Value ...interface{}) js.Value {
	return this.Element.Call("blur", Value...)
}

/*
focus()
Attaches/Triggers the focus event
*/
func (this JQueryElement) Focus(Value ...interface{}) js.Value {
	return this.Element.Call("focus", Value...)
}

/*
focusin()
Attaches an event handler to the focusin event
*/
func (this JQueryElement) FocusIn(Value ...interface{}) js.Value {
	return this.Element.Call("focusin", Value...)
}

/*
focusout()
Attaches an event handler to the focusout event
*/
func (this JQueryElement) FocusOut(Value ...interface{}) js.Value {
	return this.Element.Call("focusout", Value...)
}

/*
keydown()
Attaches/Triggers the keydown event
*/
func (this JQueryElement) KeyDown(Value ...interface{}) js.Value {
	return this.Element.Call("keydown", Value...)
}

/*
keypress()
Attaches/Triggers the keypress event
*/
func (this JQueryElement) KeyPress(Value ...interface{}) js.Value {
	return this.Element.Call("keypress", Value...)
}

/*
keyup()
Attaches/Triggers the keyup event
*/
func (this JQueryElement) KeyUp(Value ...interface{}) js.Value {
	return this.Element.Call("keyup", Value...)
}

/*
mousedown()
Attaches/Triggers the mousedown event
*/
func (this JQueryElement) MouseDown(Value ...interface{}) js.Value {
	return this.Element.Call("mousedown", Value...)
}

/*
mouseenter()
Attaches/Triggers the mouseenter event
*/
func (this JQueryElement) MouseEnter(Value ...interface{}) js.Value {
	return this.Element.Call("mouseenter", Value...)
}

/*
mouseleave()
Attaches/Triggers the mouseleave event
*/
func (this JQueryElement) MouseLeave(Value ...interface{}) js.Value {
	return this.Element.Call("mouseleave", Value...)
}

/*
mousemove()
Attaches/Triggers the mousemove event
*/
func (this JQueryElement) MouseMove(Value ...interface{}) js.Value {
	return this.Element.Call("mousemove", Value...)
}

/*
mouseout()
Attaches/Triggers the mouseout event
*/
func (this JQueryElement) MouseOut(Value ...interface{}) js.Value {
	return this.Element.Call("mouseout", Value...)
}

/*
mouseover()
Attaches/Triggers the mouseover event
*/
func (this JQueryElement) MouseOver(Value ...interface{}) js.Value {
	return this.Element.Call("mouseover", Value...)
}

/*
mouseup()
Attaches/Triggers the mouseup event
*/
func (this JQueryElement) MouseUp(Value ...interface{}) js.Value {
	return this.Element.Call("mouseup", Value...)
}

/*
resize()
Attaches/Triggers the resize event
*/
func (this JQueryElement) Resize(Value ...interface{}) js.Value {
	return this.Element.Call("resize", Value...)
}

/*
scroll()
Attaches/Triggers the scroll event
*/
func (this JQueryElement) Scroll(Value ...interface{}) js.Value {
	return this.Element.Call("scroll", Value...)
}

/*
select()
Attaches/Triggers the select event
*/
func (this JQueryElement) Select(Value ...interface{}) js.Value {
	return this.Element.Call("select", Value...)
}

/*
submit()
Attaches/Triggers the submit event
*/
func (this JQueryElement) Submit(Value ...interface{}) js.Value {
	return this.Element.Call("submit", Value...)
}

/*
ready()
Specifies a function to execute when the DOM is fully loaded
*/
func (this JQueryElement) Ready(Value interface{}) js.Value {
	return this.Element.Call("ready", Value)
}

func (this JQueryElement) Html(Value ...interface{}) js.Value {
	return this.Element.Call("html", Value...)
}

func (this JQueryElement) AppendTo(Value ...interface{}) js.Value {
	return this.Element.Call("appendTo", Value...)
}

func (this JQueryElement) AddClass(Value ...interface{}) js.Value {
	return this.Element.Call("addClass", Value...)
}

func (this JQueryElement) RemoveClass(Value ...interface{}) js.Value {
	return this.Element.Call("removeClass", Value...)
}

func (this JQueryElement) Val(Value ...interface{}) js.Value {
	return this.Element.Call("val", Value...)
}

func (this JQueryElement) Attr(Value ...interface{}) js.Value {
	return this.Element.Call("attr", Value...)
}

func (this JQueryElement) Prop(Value ...interface{}) js.Value {
	return this.Element.Call("prop", Value...)
}

/*
jQuery Event Methods
Event methods trigger or attach a function to an event handler for the selected elements.

The following table lists all the jQuery methods used to handle events.

Method / Property	Description
change()			Attaches/Triggers the change event
hover()				Attaches two event handlers to the hover event
off()				Removes event handlers attached with the on() method
on()				Attaches event handlers to elements
one()				Adds one or more event handlers to selected elements. This handler can only be triggered once per element
trigger()			Triggers all events bound to the selected elements
triggerHandler()	Triggers all functions bound to a specified event for the selected elements

*/

/*
jQuery Selectors
Use our jQuery Selector Tester to demonstrate the different selectors.

Selector	Example	Selects
*	$("*")	All elements
#id	$("#lastname")	The element with id="lastname"
.class	$(".intro")	All elements with class="intro"
.class,.class	$(".intro,.demo")	All elements with the class "intro" or "demo"
element	$("p")	All <p> elements
el1,el2,el3	$("h1,div,p")	All <h1>, <div> and <p> elements

:first	$("p:first")	The first <p> element
:last	$("p:last")	The last <p> element
:even	$("tr:even")	All even <tr> elements
:odd	$("tr:odd")	All odd <tr> elements

:first-child	$("p:first-child")	All <p> elements that are the first child of their parent
:first-of-type	$("p:first-of-type")	All <p> elements that are the first <p> element of their parent
:last-child	$("p:last-child")	All <p> elements that are the last child of their parent
:last-of-type	$("p:last-of-type")	All <p> elements that are the last <p> element of their parent
:nth-child(n)	$("p:nth-child(2)")	All <p> elements that are the 2nd child of their parent
:nth-last-child(n)	$("p:nth-last-child(2)")	All <p> elements that are the 2nd child of their parent, counting from the last child
:nth-of-type(n)	$("p:nth-of-type(2)")	All <p> elements that are the 2nd <p> element of their parent
:nth-last-of-type(n)	$("p:nth-last-of-type(2)")	All <p> elements that are the 2nd <p> element of their parent, counting from the last child
:only-child	$("p:only-child")	All <p> elements that are the only child of their parent
:only-of-type	$("p:only-of-type")	All <p> elements that are the only child, of its type, of their parent

parent > child	$("div > p")	All <p> elements that are a direct child of a <div> element
parent descendant	$("div p")	All <p> elements that are descendants of a <div> element
element + next	$("div + p")	The <p> element that are next to each <div> elements
element ~ siblings	$("div ~ p")	All <p> elements that are siblings of a <div> element

:eq(index)	$("ul li:eq(3)")	The fourth element in a list (index starts at 0)
:gt(no)	$("ul li:gt(3)")	List elements with an index greater than 3
:lt(no)	$("ul li:lt(3)")	List elements with an index less than 3
:not(selector)	$("input:not(:empty)")	All input elements that are not empty

:header	$(":header")	All header elements <h1>, <h2> ...
:animated	$(":animated")	All animated elements
:focus	$(":focus")	The element that currently has focus
:contains(text)	$(":contains('Hello')")	All elements which contains the text "Hello"
:has(selector)	$("div:has(p)")	All <div> elements that have a <p> element
:empty	$(":empty")	All elements that are empty
:parent	$(":parent")	All elements that are a parent of another element
:hidden	$("p:hidden")	All hidden <p> elements
:visible	$("table:visible")	All visible tables
:root	$(":root")	The document's root element
:lang(language)	$("p:lang(de)")	All <p> elements with a lang attribute value starting with "de"

[attribute]	$("[href]")	All elements with a href attribute
[attribute=value]	$("[href='default.htm']")	All elements with a href attribute value equal to "default.htm"
[attribute!=value]	$("[href!='default.htm']")	All elements with a href attribute value not equal to "default.htm"
[attribute$=value]	$("[href$='.jpg']")	All elements with a href attribute value ending with ".jpg"
[attribute|=value]	$("[title|='Tomorrow']")	All elements with a title attribute value equal to 'Tomorrow', or starting with 'Tomorrow' followed by a hyphen
[attribute^=value]	$("[title^='Tom']")	All elements with a title attribute value starting with "Tom"
[attribute~=value]	$("[title~='hello']")	All elements with a title attribute value containing the specific word "hello"
[attribute*=value]	$("[title*='hello']")	All elements with a title attribute value containing the word "hello"

:input	$(":input")	All input elements
:text	$(":text")	All input elements with type="text"
:password	$(":password")	All input elements with type="password"
:radio	$(":radio")	All input elements with type="radio"
:checkbox	$(":checkbox")	All input elements with type="checkbox"
:submit	$(":submit")	All input elements with type="submit"
:reset	$(":reset")	All input elements with type="reset"
:button	$(":button")	All input elements with type="button"
:image	$(":image")	All input elements with type="image"
:file	$(":file")	All input elements with type="file"
:enabled	$(":enabled")	All enabled input elements
:disabled	$(":disabled")	All disabled input elements
:selected	$(":selected")	All selected input elements
:checked	$(":checked")	All checked input elements
*/
