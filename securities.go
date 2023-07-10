package gomoex

import (
	"context"

	"github.com/tidwall/gjson"
)

// Security содержит информацию о ценной бумаге.
type Security struct {
	Ticker              string
	LotSize             int
	ISIN                string
	Board               string
	Type                string
	Instrument          string
	Shortname           string
	Prevprice           float64
	Facevalue           int64
	Status              string
	Boardname           string
	Decimals            int64
	Secname             string
	Remarks             string
	Marketcode          string
	Sectorid            string
	Minstep             float64
	Prevwaprice         float64
	Faceunit            string
	Prevdate            string
	Issuesize           int64
	Latname             string
	Regnumber           string
	Prevlegalcloseprice float64
	Currencyid          string
	Listlevel           int64
	Settledate          string
}

const (
	_securitySECID               = `SECID`
	_securityLotSize             = `LOTSIZE`
	_securityISIN                = `ISIN`
	_securityBoard               = `BOARDID`
	_securityType                = `SECTYPE`
	_securityInstrument          = `INSTRID`
	_securityShortname           = `SHORTNAME`
	_securityPrevprice           = `PREVPRICE`
	_securityFacevalue           = `FACEVALUE`
	_securityStatus              = `STATUS`
	_securityBoardname           = `BOARDNAME`
	_securityDecimals            = `DECIMALS`
	_securitySecname             = `SECNAME`
	_securityRemarks             = `REMARKS`
	_securityMarketcode          = `MARKETCODE`
	_securitySectorid            = `SECTORID`
	_securityMinstep             = `MINSTEP`
	_securityPrevwaprice         = `PREVWAPRICE`
	_securityFaceunit            = `FACEUNIT`
	_securityPrevdate            = `PREVDATE`
	_securityIssuesize           = `ISSUESIZE`
	_securityLatname             = `LATNAME`
	_securityRegnumber           = `REGNUMBER`
	_securityPrevlegalcloseprice = `PREVLEGALCLOSEPRICE`
	_securityCurrencyid          = `CURRENCYID`
	_securityListlevel           = `LISTLEVEL`
	_securitySettledate          = `SETTLEDATE`
)

type T struct {
	SECID               string      `json:"SECID"`
	LOTSIZE             int         `json:"LOTSIZE"`
	ISIN                string      `json:"ISIN"`
	BOARDID             string      `json:"BOARDID"`
	SECTYPE             string      `json:"SECTYPE"`
	INSTRID             string      `json:"INSTRID"`
	SHORTNAME           string      `json:"SHORTNAME"`
	PREVPRICE           float64     `json:"PREVPRICE"`
	FACEVALUE           int         `json:"FACEVALUE"`
	STATUS              string      `json:"STATUS"`
	BOARDNAME           string      `json:"BOARDNAME"`
	DECIMALS            int         `json:"DECIMALS"`
	SECNAME             string      `json:"SECNAME"`
	REMARKS             interface{} `json:"REMARKS"`
	MARKETCODE          string      `json:"MARKETCODE"`
	SECTORID            interface{} `json:"SECTORID"`
	MINSTEP             float64     `json:"MINSTEP"`
	PREVWAPRICE         float64     `json:"PREVWAPRICE"`
	FACEUNIT            string      `json:"FACEUNIT"`
	PREVDATE            string      `json:"PREVDATE"`
	ISSUESIZE           int         `json:"ISSUESIZE"`
	LATNAME             string      `json:"LATNAME"`
	REGNUMBER           string      `json:"REGNUMBER"`
	PREVLEGALCLOSEPRICE float64     `json:"PREVLEGALCLOSEPRICE"`
	CURRENCYID          string      `json:"CURRENCYID"`
	LISTLEVEL           int         `json:"LISTLEVEL"`
	SETTLEDATE          string      `json:"SETTLEDATE"`
}

func securityConverter(row gjson.Result) (interface{}, error) {
	var sec Security

	sec.Ticker = row.Get(_securitySECID).String()
	sec.LotSize = int(row.Get(_securityLotSize).Int())
	sec.ISIN = row.Get(_securityISIN).String()
	sec.Board = row.Get(_securityBoard).String()
	sec.Type = row.Get(_securityType).String()
	sec.Instrument = row.Get(_securityInstrument).String()
	sec.Shortname = row.Get(_securityShortname).String()
	sec.Prevprice = row.Get(_securityPrevprice).Float()
	sec.Facevalue = row.Get(_securityFacevalue).Int()
	sec.Status = row.Get(_securityStatus).String()
	sec.Boardname = row.Get(_securityBoardname).String()
	sec.Decimals = row.Get(_securityDecimals).Int()
	sec.Secname = row.Get(_securitySecname).String()
	sec.Remarks = row.Get(_securityRemarks).String()
	sec.Marketcode = row.Get(_securityMarketcode).String()
	sec.Sectorid = row.Get(_securitySectorid).String()
	sec.Minstep = row.Get(_securityMinstep).Float()
	sec.Prevwaprice = row.Get(_securityPrevwaprice).Float()
	sec.Faceunit = row.Get(_securityFaceunit).String()
	sec.Prevdate = row.Get(_securityPrevdate).String()
	sec.Issuesize = row.Get(_securityIssuesize).Int()
	sec.Latname = row.Get(_securityLatname).String()
	sec.Regnumber = row.Get(_securityRegnumber).String()
	sec.Prevlegalcloseprice = row.Get(_securityPrevlegalcloseprice).Float()
	sec.Currencyid = row.Get(_securityCurrencyid).String()
	sec.Listlevel = row.Get(_securityListlevel).Int()
	sec.Settledate = row.Get(_securitySettledate).String()

	return sec, nil
}

// BoardSecurities получает таблицу с торгуемыми бумагами в данном режиме торгов.
//
// Описание запроса - https://iss.moex.com/iss/reference/32
func (iss *ISSClient) BoardSecurities(ctx context.Context, engine, market, board string) (table []Security, err error) {
	query := querySettings{
		engine:       engine,
		market:       market,
		board:        board,
		object:       "securities",
		table:        "securities",
		rowConverter: securityConverter,
	}

	for raw := range iss.rowsGen(ctx, query.Make()) {
		switch row := raw.(type) {
		case Security:
			table = append(table, row)
		case error:
			return nil, row
		}
	}

	return table, nil
}
