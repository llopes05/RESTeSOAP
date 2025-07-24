package soap

import "encoding/xml"

type EmprestarLivroEnvelope struct {
	XMLName xml.Name           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Body    EmprestarLivroBody `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`
}

type EmprestarLivroBody struct {
	XMLName xml.Name              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`
	Request EmprestarLivroRequest `xml:"EmprestarLivroRequest"`
}

type EmprestarLivroRequest struct {
	XMLName   xml.Name `xml:"http://biblioteca.com/soap/definitions EmprestarLivroRequest"`
	UsuarioID uint     `xml:"usuario_id"`
	LivroID   uint     `xml:"livro_id"`
}


type SOAPResponseEnvelope struct {
	XMLName xml.Name         `xml:"soap:Envelope"`
	Body    SOAPResponseBody `xml:"soap:Body"`
}

type SOAPResponseBody struct {
	XMLName  xml.Name    `xml:"soap:Body"`
	Response interface{} `xml:",any"` 
}

type EmprestarLivroResponse struct {
	XMLName      xml.Name `xml:"tns:EmprestarLivroResponse"`
	Sucesso      bool     `xml:"sucesso"`
	Mensagem     string   `xml:"mensagem"`
	EmprestimoID uint     `xml:"emprestimo_id,omitempty"`
}

type SOAPFault struct {
	XMLName     xml.Name `xml:"soap:Fault"`
	FaultCode   string   `xml:"faultcode"`
	FaultString string   `xml:"faultstring"`
}