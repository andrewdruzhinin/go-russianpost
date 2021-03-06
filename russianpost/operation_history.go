package russianpost

import (
	"encoding/xml"
	"fmt"
	"strings"
	"time"
)

// SoapRequest структура soap запроса
type SoapRequest struct {
	XMLName     xml.Name `xml:"soap:Envelope"`
	SoapAttr    string   `xml:"xmlns:soap,attr"`
	OperAttr    string   `xml:"xmlns:oper,attr"`
	DataAttr    string   `xml:"xmlns:data,attr"`
	SoapEnvAttr string   `xml:"xmlns:soapenv,attr"`
	Header      string   `xml:"soap:Header"`
	Body        SoapRequestBody
}

// SoapRequestBody тело запроса
type SoapRequestBody struct {
	XMLName xml.Name `xml:"soap:Body"`
	Oper    SoapRequestOper
}

// SoapRequestOper вид операции, передаваемой в запросе
type SoapRequestOper struct {
	XMLName     xml.Name `xml:"oper:getOperationHistory"`
	OperHistReq OperationHistoryRequest
	AuthHeader  AuthorizationHeader
}

// OperationHistoryRequest Содержит элементы Barcode, MessageType, Language
type OperationHistoryRequest struct {
	XMLName     xml.Name `xml:"data:OperationHistoryRequest"`
	Barcode     string   `xml:"data:Barcode"`
	MessageType string   `xml:"data:MessageType"`
	Language    string   `xml:"data:Language"`
}

// AuthorizationHeader Содержит элементы login и password.
// Атрибут soapenv:mustUnderstand элемента AuthorizationHeader должен содержать значение 1.
type AuthorizationHeader struct {
	XMLName        xml.Name `xml:"data:AuthorizationHeader"`
	MustUnderstand string   `xml:"soapenv:mustUnderstand,attr"`
	Login          string   `xml:"data:login"`
	Password       string   `xml:"data:password"`
}

//Result результат запроса
type Result struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    Body     `xml:"Body"`
}

// Body тело ответа
type Body struct {
	GetOperationHistoryResponse GetOperationHistoryResponse `xml:"getOperationHistoryResponse"`
}

// GetOperationHistoryResponse элемент xml структуры
type GetOperationHistoryResponse struct {
	OperationHistoryData OperationHistoryData `xml:"OperationHistoryData"`
}

// OperationHistoryData содержит список элементов historyRecord. Каждый из них содержит информацию об одной операции над отправлением
// Если над отправлением еще не зарегистрировано ни одной операции, то возвращается пустой список элементов historyRecord
type OperationHistoryData struct {
	HistoryRecords []HistoryRecord `xml:"historyRecord"`
}

// HistoryRecord содержит информацию об одной операции над отправлением.
type HistoryRecord struct {
	AddressParameters   AddressParameters   `xml:"AddressParameters"`
	FinanceParameters   FinanceParameters   `xml:"FinanceParameters"`
	ItemParameters      ItemParameters      `xml:"ItemParameters"`
	OperationParameters OperationParameters `xml:"OperationParameters"`
	UserParameters      UserParameters      `xml:"UserParameters"`
}

// UserParameters Содержит данные субъектов, связанных с операцией над почтовым отправлением
type UserParameters struct {
	SendCtg SendCtg `xml:"SendCtg"`
	Sndr    string
	Rcpn    string
}

// SendCtg Содержит информацию о категории отправителя.
type SendCtg struct {
	ID   int `xml:"Id"`
	Name string
}

// OperationParameters Cодержит параметры операции над отправлением
type OperationParameters struct {
	OperType OperType `xml:"OperType"`
	OperAttr OperAttr `xml:"OperAttr"`
	OperDate string
}

//OperAttr Содержит информацию об атрибуте операции над отправлением
type OperAttr struct {
	ID   int `xml:"Id"`
	Name string
}

// OperType Содержит информацию об операции над отправлением
type OperType struct {
	ID   int `xml:"Id"`
	Name string
}

// ItemParameters Содержит данные о почтовом отправлении
type ItemParameters struct {
	Barcode         string
	Internum        string
	ValidRuType     string
	ValidEnType     string
	ComplexItemName string
	MailRank        MailRank `xml:"MailRank"`
	MailType        MailType `xml:"MailType"`
	Mass            int
	MaxMassRu       int
	MaxMassEn       int
}

// MailRank Содержит информацию о разряде почтового отправления
type MailRank struct {
	ID   int `xml:"Id"`
	Name string
}

// MailType Содержит данные о виде почтового отправления
type MailType struct {
	ID   int `xml:"Id"`
	Name string
}

// MailCtg Содержит данные о категории почтового отправления.
type MailCtg struct {
	ID   int `xml:"Id"`
	Name string
}

// FinanceParameters Содержит финансовые данные, связанные с операцией над почтовым отправлением
type FinanceParameters struct {
	Payment    int
	Value      int
	MassRate   int
	InsrRate   int
	AirRate    int
	Rate       int
	CustomDuty int
}

// AddressParameters Содержит адресные данные с операцией над отправлением
type AddressParameters struct {
	DestinationAddress DestinationAddress `xml:"DestinationAddress"`
	OperationAddress   OperationAddress   `xml:"OperationAddress"`
	MailDirect         MailDirect         `xml:"MailDirect"`
	CountryOper        CountryOper        `xml:"CountryOper"`
	CountryFrom        CountryFrom        `xml:"CountryFrom"`
}

// CountryFrom Содержит данные о стране приема почтового отправления
type CountryFrom struct {
	ID     int `xml:"Id"`
	Code2A string
	Code3A string
	NameRU string
	NameEN string
}

// CountryOper Содержит данные о стране проведения операции над почтовым отправлением
type CountryOper struct {
	ID     int `xml:"Id"`
	Code2A string
	Code3A string
	NameRU string
	NameEN string
}

// MailDirect Содержит данные о стране места назначения пересылки отправления
type MailDirect struct {
	ID     int `xml:"Id"`
	Code2A string
	Code3A string
	NameRU string
	NameEN string
}

// OperationAddress Содержит адресные данные места проведения операции над отправлением
type OperationAddress struct {
	Index       string `xml:"Index"`
	Description string `xml:"Description"`
}

// DestinationAddress Содержит адресные данные места назначения пересылки отправления
type DestinationAddress struct {
	Index       string `xml:"Index"`
	Description string `xml:"Description"`
}

// Data содержит данные об операциях над почтовым отправлением
type Data struct {
	DataItems []DataItem
}

// DataItem содержит данные об операции над почтовым отправлением
type DataItem struct {
	DestinationAddress string // место назначения
	Mass               int    // вес отправления
	OperarationDate    string // дата проведения оперрации над отправлением
	Operation          string // операция над отправлением (тип и атрибут)
	OperationLocation  string // место проведения операции над отправлением (почтовый индекс, название)
}

// GetOperationHistory возвращает историю операций над отправлением
// barcode - Идентификатор регистрируемого почтового отправления в одном из форматов:
//	- внутрироссийский, состоящий из 14 символов (цифровой);
//	- международный, состоящий из 13 символов (буквенно-цифровой) в формате S10.
// messegeType - Тип сообщения. Возможные значения:
//	0 - история операций для отправления;
//	1 - история операций для заказного уведомления по данному отправлению.
// language - Язык, на котором должны возвращаться названия операций/атрибутов и сообщения об ошибках. Допустимые значения:
//	RUS – использовать русский язык (используется по умолчанию);
//	ENG – использовать английский язык.
func (c *Client) GetOperationHistory(barcode, messegeType, language string) (Data, error) {
	operHistReq := OperationHistoryRequest{Barcode: barcode, MessageType: messegeType, Language: language}
	authHeader := AuthorizationHeader{MustUnderstand: "1", Login: c.login, Password: c.password}
	soapRequestOper := SoapRequestOper{OperHistReq: operHistReq, AuthHeader: authHeader}
	soapRequestBody := SoapRequestBody{Oper: soapRequestOper}
	soapRequest := SoapRequest{
		SoapAttr:    "http://www.w3.org/2003/05/soap-envelope",
		OperAttr:    "http://russianpost.org/operationhistory",
		DataAttr:    "http://russianpost.org/operationhistory/data",
		SoapEnvAttr: "http://schemas.xmlsoap.org/soap/envelope/",
		Body:        soapRequestBody,
	}
	var data Data
	xmlSoapRequest, err := xml.MarshalIndent(soapRequest, "", "    ")
	if err != nil {
		fmt.Println(err)
		return data, err
	}
	payload := strings.NewReader(string(xmlSoapRequest))
	req, _ := c.NewRequest("POST", "", payload)

	req.Header.Add("content-type", "application/soap+xml")

	body, err := c.Do(req)
	if err != nil {
		return data, err
	}
	result := Result{}
	err = xml.Unmarshal([]byte(body), &result)
	if err != nil {
		fmt.Printf("error: %v", err)
		return data, err
	}
	data = buildData(result)
	return data, nil
}

// buildData преобразовываем данные об операциях над отправлением в нужный нам вид Data/DataItem
func buildData(result Result) Data {
	var data Data

	for _, historyRecord := range result.Body.GetOperationHistoryResponse.OperationHistoryData.HistoryRecords {
		var dataItem DataItem
		dataItem.DestinationAddress = historyRecord.AddressParameters.DestinationAddress.Index + " " + historyRecord.AddressParameters.DestinationAddress.Description
		dataItem.Mass = historyRecord.ItemParameters.Mass
		dataItem.OperarationDate = dateFormat(historyRecord.OperationParameters.OperDate)
		dataItem.Operation = historyRecord.OperationParameters.OperType.Name + " " + historyRecord.OperationParameters.OperAttr.Name
		dataItem.OperationLocation = historyRecord.AddressParameters.OperationAddress.Description + " " + historyRecord.AddressParameters.OperationAddress.Index
		data.DataItems = append(data.DataItems, dataItem)
	}

	return data
}

// dateFormat преобразовывем дату в нужный формат, возвращаем в виде строки
func dateFormat(dateTime string) string {
	t, err := time.Parse(time.RFC3339, dateTime)
	if err != nil {
		fmt.Println(err)
	}
	dateTimeResult := fmt.Sprintf("%d %s %d %d:%d", t.Day(), t.Month(), t.Year(), t.Hour(), t.Minute())
	return dateTimeResult
}
