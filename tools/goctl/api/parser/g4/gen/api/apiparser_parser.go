// Code generated from C:/Users/Psydu/GolandProjects/go-zero/tools/goctl/api/parser/g4\ApiParser.g4 by ANTLR 4.12.0. DO NOT EDIT.

package api // ApiParser
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type ApiParserParser struct {
	*antlr.BaseParser
}

var apiparserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func apiparserParserInit() {
	staticData := &apiparserParserStaticData
	staticData.literalNames = []string{
		"", "'='", "'('", "')'", "'{'", "'}'", "'*'", "'time.Time'", "'['",
		"']'", "'returns'", "'-'", "'/'", "'/:'", "'@doc'", "'@handler'", "'@desc'",
		"'interface{}'", "'@server'", "'@type'", "'<-'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "ATDOC", "ATHANDLER",
		"ATDESC", "INTERFACE", "ATSERVER", "ATTYPE", "AS", "WS", "COMMENT",
		"LINE_COMMENT", "STRING", "RAW_STRING", "LINE_VALUE", "ID", "LetterOrDigit",
	}
	staticData.ruleNames = []string{
		"api", "spec", "syntaxLit", "importSpec", "importLit", "importBlock",
		"importBlockValue", "importValue", "infoSpec", "svcommLit", "typeSpec",
		"typeLit", "typeBlock", "typeLitBody", "typeBlockBody", "atType", "typeStruct",
		"typeAlias", "typeBlockStruct", "typeBlockAlias", "field", "normalField",
		"anonymousFiled", "dataType", "pointerType", "mapType", "arrayType",
		"serviceSpec", "atServer", "serviceApi", "serviceRoute", "atDoc", "atHandler",
		"atDesc", "route", "body", "replybody", "kvLit", "serviceName", "path",
		"pathItem",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 28, 392, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 1, 0, 5, 0, 84,
		8, 0, 10, 0, 12, 0, 87, 9, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1,
		95, 8, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 3, 3, 105, 8,
		3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 4, 5, 115, 8, 5, 11,
		5, 12, 5, 116, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1,
		8, 1, 8, 4, 8, 130, 8, 8, 11, 8, 12, 8, 131, 1, 8, 1, 8, 1, 9, 1, 9, 1,
		9, 1, 9, 1, 10, 1, 10, 3, 10, 142, 8, 10, 1, 11, 3, 11, 145, 8, 11, 1,
		11, 1, 11, 1, 11, 1, 11, 1, 12, 3, 12, 152, 8, 12, 1, 12, 1, 12, 1, 12,
		1, 12, 5, 12, 158, 8, 12, 10, 12, 12, 12, 161, 9, 12, 1, 12, 1, 12, 1,
		13, 1, 13, 3, 13, 167, 8, 13, 1, 14, 1, 14, 3, 14, 171, 8, 14, 1, 15, 1,
		15, 3, 15, 175, 8, 15, 1, 15, 4, 15, 178, 8, 15, 11, 15, 12, 15, 179, 1,
		15, 3, 15, 183, 8, 15, 1, 15, 3, 15, 186, 8, 15, 1, 16, 1, 16, 1, 16, 3,
		16, 191, 8, 16, 1, 16, 1, 16, 5, 16, 195, 8, 16, 10, 16, 12, 16, 198, 9,
		16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 3, 17, 205, 8, 17, 1, 17, 1, 17,
		1, 18, 1, 18, 1, 18, 3, 18, 212, 8, 18, 1, 18, 1, 18, 5, 18, 216, 8, 18,
		10, 18, 12, 18, 219, 9, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 3, 19, 226,
		8, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 3, 20, 233, 8, 20, 1, 21, 1,
		21, 1, 21, 1, 21, 3, 21, 239, 8, 21, 1, 22, 3, 22, 242, 8, 22, 1, 22, 1,
		22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 3, 23, 254,
		8, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1,
		25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 26, 1, 27, 3, 27, 273, 8, 27,
		1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 4, 28, 280, 8, 28, 11, 28, 12, 28, 281,
		1, 28, 1, 28, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 5, 29, 291, 8, 29, 10,
		29, 12, 29, 294, 9, 29, 1, 29, 1, 29, 1, 30, 3, 30, 299, 8, 30, 1, 30,
		1, 30, 3, 30, 303, 8, 30, 1, 30, 3, 30, 306, 8, 30, 1, 30, 1, 30, 1, 31,
		1, 31, 3, 31, 312, 8, 31, 1, 31, 4, 31, 315, 8, 31, 11, 31, 12, 31, 316,
		1, 31, 3, 31, 320, 8, 31, 1, 31, 3, 31, 323, 8, 31, 1, 32, 1, 32, 1, 32,
		1, 33, 1, 33, 1, 33, 1, 34, 1, 34, 1, 34, 1, 34, 3, 34, 335, 8, 34, 1,
		34, 3, 34, 338, 8, 34, 1, 35, 1, 35, 3, 35, 342, 8, 35, 1, 35, 1, 35, 1,
		36, 1, 36, 1, 36, 3, 36, 349, 8, 36, 1, 36, 1, 36, 1, 37, 1, 37, 1, 37,
		1, 37, 1, 38, 1, 38, 3, 38, 359, 8, 38, 4, 38, 361, 8, 38, 11, 38, 12,
		38, 362, 1, 39, 1, 39, 1, 39, 1, 39, 5, 39, 369, 8, 39, 10, 39, 12, 39,
		372, 9, 39, 1, 39, 1, 39, 1, 39, 1, 39, 3, 39, 378, 8, 39, 4, 39, 380,
		8, 39, 11, 39, 12, 39, 381, 1, 39, 3, 39, 385, 8, 39, 1, 40, 4, 40, 388,
		8, 40, 11, 40, 12, 40, 389, 1, 40, 0, 0, 41, 0, 2, 4, 6, 8, 10, 12, 14,
		16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50,
		52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 0, 1, 1, 0,
		27, 28, 406, 0, 85, 1, 0, 0, 0, 2, 94, 1, 0, 0, 0, 4, 96, 1, 0, 0, 0, 6,
		104, 1, 0, 0, 0, 8, 106, 1, 0, 0, 0, 10, 110, 1, 0, 0, 0, 12, 120, 1, 0,
		0, 0, 14, 122, 1, 0, 0, 0, 16, 125, 1, 0, 0, 0, 18, 135, 1, 0, 0, 0, 20,
		141, 1, 0, 0, 0, 22, 144, 1, 0, 0, 0, 24, 151, 1, 0, 0, 0, 26, 166, 1,
		0, 0, 0, 28, 170, 1, 0, 0, 0, 30, 172, 1, 0, 0, 0, 32, 187, 1, 0, 0, 0,
		34, 201, 1, 0, 0, 0, 36, 208, 1, 0, 0, 0, 38, 222, 1, 0, 0, 0, 40, 232,
		1, 0, 0, 0, 42, 234, 1, 0, 0, 0, 44, 241, 1, 0, 0, 0, 46, 253, 1, 0, 0,
		0, 48, 255, 1, 0, 0, 0, 50, 259, 1, 0, 0, 0, 52, 267, 1, 0, 0, 0, 54, 272,
		1, 0, 0, 0, 56, 276, 1, 0, 0, 0, 58, 285, 1, 0, 0, 0, 60, 298, 1, 0, 0,
		0, 62, 309, 1, 0, 0, 0, 64, 324, 1, 0, 0, 0, 66, 327, 1, 0, 0, 0, 68, 330,
		1, 0, 0, 0, 70, 339, 1, 0, 0, 0, 72, 345, 1, 0, 0, 0, 74, 352, 1, 0, 0,
		0, 76, 360, 1, 0, 0, 0, 78, 384, 1, 0, 0, 0, 80, 387, 1, 0, 0, 0, 82, 84,
		3, 2, 1, 0, 83, 82, 1, 0, 0, 0, 84, 87, 1, 0, 0, 0, 85, 83, 1, 0, 0, 0,
		85, 86, 1, 0, 0, 0, 86, 1, 1, 0, 0, 0, 87, 85, 1, 0, 0, 0, 88, 95, 3, 4,
		2, 0, 89, 95, 3, 18, 9, 0, 90, 95, 3, 6, 3, 0, 91, 95, 3, 16, 8, 0, 92,
		95, 3, 20, 10, 0, 93, 95, 3, 54, 27, 0, 94, 88, 1, 0, 0, 0, 94, 89, 1,
		0, 0, 0, 94, 90, 1, 0, 0, 0, 94, 91, 1, 0, 0, 0, 94, 92, 1, 0, 0, 0, 94,
		93, 1, 0, 0, 0, 95, 3, 1, 0, 0, 0, 96, 97, 6, 2, -1, 0, 97, 98, 5, 27,
		0, 0, 98, 99, 5, 1, 0, 0, 99, 100, 6, 2, -1, 0, 100, 101, 5, 24, 0, 0,
		101, 5, 1, 0, 0, 0, 102, 105, 3, 8, 4, 0, 103, 105, 3, 10, 5, 0, 104, 102,
		1, 0, 0, 0, 104, 103, 1, 0, 0, 0, 105, 7, 1, 0, 0, 0, 106, 107, 6, 4, -1,
		0, 107, 108, 5, 27, 0, 0, 108, 109, 3, 14, 7, 0, 109, 9, 1, 0, 0, 0, 110,
		111, 6, 5, -1, 0, 111, 112, 5, 27, 0, 0, 112, 114, 5, 2, 0, 0, 113, 115,
		3, 12, 6, 0, 114, 113, 1, 0, 0, 0, 115, 116, 1, 0, 0, 0, 116, 114, 1, 0,
		0, 0, 116, 117, 1, 0, 0, 0, 117, 118, 1, 0, 0, 0, 118, 119, 5, 3, 0, 0,
		119, 11, 1, 0, 0, 0, 120, 121, 3, 14, 7, 0, 121, 13, 1, 0, 0, 0, 122, 123,
		6, 7, -1, 0, 123, 124, 5, 24, 0, 0, 124, 15, 1, 0, 0, 0, 125, 126, 6, 8,
		-1, 0, 126, 127, 5, 27, 0, 0, 127, 129, 5, 2, 0, 0, 128, 130, 3, 74, 37,
		0, 129, 128, 1, 0, 0, 0, 130, 131, 1, 0, 0, 0, 131, 129, 1, 0, 0, 0, 131,
		132, 1, 0, 0, 0, 132, 133, 1, 0, 0, 0, 133, 134, 5, 3, 0, 0, 134, 17, 1,
		0, 0, 0, 135, 136, 5, 27, 0, 0, 136, 137, 5, 20, 0, 0, 137, 138, 5, 24,
		0, 0, 138, 19, 1, 0, 0, 0, 139, 142, 3, 22, 11, 0, 140, 142, 3, 24, 12,
		0, 141, 139, 1, 0, 0, 0, 141, 140, 1, 0, 0, 0, 142, 21, 1, 0, 0, 0, 143,
		145, 3, 30, 15, 0, 144, 143, 1, 0, 0, 0, 144, 145, 1, 0, 0, 0, 145, 146,
		1, 0, 0, 0, 146, 147, 6, 11, -1, 0, 147, 148, 5, 27, 0, 0, 148, 149, 3,
		26, 13, 0, 149, 23, 1, 0, 0, 0, 150, 152, 3, 30, 15, 0, 151, 150, 1, 0,
		0, 0, 151, 152, 1, 0, 0, 0, 152, 153, 1, 0, 0, 0, 153, 154, 6, 12, -1,
		0, 154, 155, 5, 27, 0, 0, 155, 159, 5, 2, 0, 0, 156, 158, 3, 28, 14, 0,
		157, 156, 1, 0, 0, 0, 158, 161, 1, 0, 0, 0, 159, 157, 1, 0, 0, 0, 159,
		160, 1, 0, 0, 0, 160, 162, 1, 0, 0, 0, 161, 159, 1, 0, 0, 0, 162, 163,
		5, 3, 0, 0, 163, 25, 1, 0, 0, 0, 164, 167, 3, 32, 16, 0, 165, 167, 3, 34,
		17, 0, 166, 164, 1, 0, 0, 0, 166, 165, 1, 0, 0, 0, 167, 27, 1, 0, 0, 0,
		168, 171, 3, 36, 18, 0, 169, 171, 3, 38, 19, 0, 170, 168, 1, 0, 0, 0, 170,
		169, 1, 0, 0, 0, 171, 29, 1, 0, 0, 0, 172, 174, 5, 19, 0, 0, 173, 175,
		5, 2, 0, 0, 174, 173, 1, 0, 0, 0, 174, 175, 1, 0, 0, 0, 175, 182, 1, 0,
		0, 0, 176, 178, 3, 74, 37, 0, 177, 176, 1, 0, 0, 0, 178, 179, 1, 0, 0,
		0, 179, 177, 1, 0, 0, 0, 179, 180, 1, 0, 0, 0, 180, 183, 1, 0, 0, 0, 181,
		183, 5, 24, 0, 0, 182, 177, 1, 0, 0, 0, 182, 181, 1, 0, 0, 0, 183, 185,
		1, 0, 0, 0, 184, 186, 5, 3, 0, 0, 185, 184, 1, 0, 0, 0, 185, 186, 1, 0,
		0, 0, 186, 31, 1, 0, 0, 0, 187, 188, 6, 16, -1, 0, 188, 190, 5, 27, 0,
		0, 189, 191, 5, 27, 0, 0, 190, 189, 1, 0, 0, 0, 190, 191, 1, 0, 0, 0, 191,
		192, 1, 0, 0, 0, 192, 196, 5, 4, 0, 0, 193, 195, 3, 40, 20, 0, 194, 193,
		1, 0, 0, 0, 195, 198, 1, 0, 0, 0, 196, 194, 1, 0, 0, 0, 196, 197, 1, 0,
		0, 0, 197, 199, 1, 0, 0, 0, 198, 196, 1, 0, 0, 0, 199, 200, 5, 5, 0, 0,
		200, 33, 1, 0, 0, 0, 201, 202, 6, 17, -1, 0, 202, 204, 5, 27, 0, 0, 203,
		205, 5, 1, 0, 0, 204, 203, 1, 0, 0, 0, 204, 205, 1, 0, 0, 0, 205, 206,
		1, 0, 0, 0, 206, 207, 3, 46, 23, 0, 207, 35, 1, 0, 0, 0, 208, 209, 6, 18,
		-1, 0, 209, 211, 5, 27, 0, 0, 210, 212, 5, 27, 0, 0, 211, 210, 1, 0, 0,
		0, 211, 212, 1, 0, 0, 0, 212, 213, 1, 0, 0, 0, 213, 217, 5, 4, 0, 0, 214,
		216, 3, 40, 20, 0, 215, 214, 1, 0, 0, 0, 216, 219, 1, 0, 0, 0, 217, 215,
		1, 0, 0, 0, 217, 218, 1, 0, 0, 0, 218, 220, 1, 0, 0, 0, 219, 217, 1, 0,
		0, 0, 220, 221, 5, 5, 0, 0, 221, 37, 1, 0, 0, 0, 222, 223, 6, 19, -1, 0,
		223, 225, 5, 27, 0, 0, 224, 226, 5, 1, 0, 0, 225, 224, 1, 0, 0, 0, 225,
		226, 1, 0, 0, 0, 226, 227, 1, 0, 0, 0, 227, 228, 3, 46, 23, 0, 228, 39,
		1, 0, 0, 0, 229, 230, 4, 20, 0, 0, 230, 233, 3, 42, 21, 0, 231, 233, 3,
		44, 22, 0, 232, 229, 1, 0, 0, 0, 232, 231, 1, 0, 0, 0, 233, 41, 1, 0, 0,
		0, 234, 235, 6, 21, -1, 0, 235, 236, 5, 27, 0, 0, 236, 238, 3, 46, 23,
		0, 237, 239, 5, 25, 0, 0, 238, 237, 1, 0, 0, 0, 238, 239, 1, 0, 0, 0, 239,
		43, 1, 0, 0, 0, 240, 242, 5, 6, 0, 0, 241, 240, 1, 0, 0, 0, 241, 242, 1,
		0, 0, 0, 242, 243, 1, 0, 0, 0, 243, 244, 5, 27, 0, 0, 244, 45, 1, 0, 0,
		0, 245, 246, 6, 23, -1, 0, 246, 254, 5, 27, 0, 0, 247, 254, 3, 50, 25,
		0, 248, 254, 3, 52, 26, 0, 249, 254, 5, 17, 0, 0, 250, 254, 5, 7, 0, 0,
		251, 254, 3, 48, 24, 0, 252, 254, 3, 32, 16, 0, 253, 245, 1, 0, 0, 0, 253,
		247, 1, 0, 0, 0, 253, 248, 1, 0, 0, 0, 253, 249, 1, 0, 0, 0, 253, 250,
		1, 0, 0, 0, 253, 251, 1, 0, 0, 0, 253, 252, 1, 0, 0, 0, 254, 47, 1, 0,
		0, 0, 255, 256, 5, 6, 0, 0, 256, 257, 6, 24, -1, 0, 257, 258, 5, 27, 0,
		0, 258, 49, 1, 0, 0, 0, 259, 260, 6, 25, -1, 0, 260, 261, 5, 27, 0, 0,
		261, 262, 5, 8, 0, 0, 262, 263, 6, 25, -1, 0, 263, 264, 5, 27, 0, 0, 264,
		265, 5, 9, 0, 0, 265, 266, 3, 46, 23, 0, 266, 51, 1, 0, 0, 0, 267, 268,
		5, 8, 0, 0, 268, 269, 5, 9, 0, 0, 269, 270, 3, 46, 23, 0, 270, 53, 1, 0,
		0, 0, 271, 273, 3, 56, 28, 0, 272, 271, 1, 0, 0, 0, 272, 273, 1, 0, 0,
		0, 273, 274, 1, 0, 0, 0, 274, 275, 3, 58, 29, 0, 275, 55, 1, 0, 0, 0, 276,
		277, 5, 18, 0, 0, 277, 279, 5, 2, 0, 0, 278, 280, 3, 74, 37, 0, 279, 278,
		1, 0, 0, 0, 280, 281, 1, 0, 0, 0, 281, 279, 1, 0, 0, 0, 281, 282, 1, 0,
		0, 0, 282, 283, 1, 0, 0, 0, 283, 284, 5, 3, 0, 0, 284, 57, 1, 0, 0, 0,
		285, 286, 6, 29, -1, 0, 286, 287, 5, 27, 0, 0, 287, 288, 3, 76, 38, 0,
		288, 292, 5, 4, 0, 0, 289, 291, 3, 60, 30, 0, 290, 289, 1, 0, 0, 0, 291,
		294, 1, 0, 0, 0, 292, 290, 1, 0, 0, 0, 292, 293, 1, 0, 0, 0, 293, 295,
		1, 0, 0, 0, 294, 292, 1, 0, 0, 0, 295, 296, 5, 5, 0, 0, 296, 59, 1, 0,
		0, 0, 297, 299, 3, 62, 31, 0, 298, 297, 1, 0, 0, 0, 298, 299, 1, 0, 0,
		0, 299, 302, 1, 0, 0, 0, 300, 303, 3, 56, 28, 0, 301, 303, 3, 64, 32, 0,
		302, 300, 1, 0, 0, 0, 302, 301, 1, 0, 0, 0, 303, 305, 1, 0, 0, 0, 304,
		306, 3, 66, 33, 0, 305, 304, 1, 0, 0, 0, 305, 306, 1, 0, 0, 0, 306, 307,
		1, 0, 0, 0, 307, 308, 3, 68, 34, 0, 308, 61, 1, 0, 0, 0, 309, 311, 5, 14,
		0, 0, 310, 312, 5, 2, 0, 0, 311, 310, 1, 0, 0, 0, 311, 312, 1, 0, 0, 0,
		312, 319, 1, 0, 0, 0, 313, 315, 3, 74, 37, 0, 314, 313, 1, 0, 0, 0, 315,
		316, 1, 0, 0, 0, 316, 314, 1, 0, 0, 0, 316, 317, 1, 0, 0, 0, 317, 320,
		1, 0, 0, 0, 318, 320, 5, 24, 0, 0, 319, 314, 1, 0, 0, 0, 319, 318, 1, 0,
		0, 0, 320, 322, 1, 0, 0, 0, 321, 323, 5, 3, 0, 0, 322, 321, 1, 0, 0, 0,
		322, 323, 1, 0, 0, 0, 323, 63, 1, 0, 0, 0, 324, 325, 5, 15, 0, 0, 325,
		326, 5, 27, 0, 0, 326, 65, 1, 0, 0, 0, 327, 328, 5, 16, 0, 0, 328, 329,
		5, 27, 0, 0, 329, 67, 1, 0, 0, 0, 330, 331, 6, 34, -1, 0, 331, 332, 5,
		27, 0, 0, 332, 334, 3, 78, 39, 0, 333, 335, 3, 70, 35, 0, 334, 333, 1,
		0, 0, 0, 334, 335, 1, 0, 0, 0, 335, 337, 1, 0, 0, 0, 336, 338, 3, 72, 36,
		0, 337, 336, 1, 0, 0, 0, 337, 338, 1, 0, 0, 0, 338, 69, 1, 0, 0, 0, 339,
		341, 5, 2, 0, 0, 340, 342, 3, 46, 23, 0, 341, 340, 1, 0, 0, 0, 341, 342,
		1, 0, 0, 0, 342, 343, 1, 0, 0, 0, 343, 344, 5, 3, 0, 0, 344, 71, 1, 0,
		0, 0, 345, 346, 5, 10, 0, 0, 346, 348, 5, 2, 0, 0, 347, 349, 3, 46, 23,
		0, 348, 347, 1, 0, 0, 0, 348, 349, 1, 0, 0, 0, 349, 350, 1, 0, 0, 0, 350,
		351, 5, 3, 0, 0, 351, 73, 1, 0, 0, 0, 352, 353, 5, 27, 0, 0, 353, 354,
		6, 37, -1, 0, 354, 355, 5, 26, 0, 0, 355, 75, 1, 0, 0, 0, 356, 358, 5,
		27, 0, 0, 357, 359, 5, 11, 0, 0, 358, 357, 1, 0, 0, 0, 358, 359, 1, 0,
		0, 0, 359, 361, 1, 0, 0, 0, 360, 356, 1, 0, 0, 0, 361, 362, 1, 0, 0, 0,
		362, 360, 1, 0, 0, 0, 362, 363, 1, 0, 0, 0, 363, 77, 1, 0, 0, 0, 364, 365,
		5, 12, 0, 0, 365, 370, 3, 80, 40, 0, 366, 367, 5, 11, 0, 0, 367, 369, 3,
		80, 40, 0, 368, 366, 1, 0, 0, 0, 369, 372, 1, 0, 0, 0, 370, 368, 1, 0,
		0, 0, 370, 371, 1, 0, 0, 0, 371, 380, 1, 0, 0, 0, 372, 370, 1, 0, 0, 0,
		373, 374, 5, 13, 0, 0, 374, 377, 3, 80, 40, 0, 375, 376, 5, 11, 0, 0, 376,
		378, 3, 80, 40, 0, 377, 375, 1, 0, 0, 0, 377, 378, 1, 0, 0, 0, 378, 380,
		1, 0, 0, 0, 379, 364, 1, 0, 0, 0, 379, 373, 1, 0, 0, 0, 380, 381, 1, 0,
		0, 0, 381, 379, 1, 0, 0, 0, 381, 382, 1, 0, 0, 0, 382, 385, 1, 0, 0, 0,
		383, 385, 5, 12, 0, 0, 384, 379, 1, 0, 0, 0, 384, 383, 1, 0, 0, 0, 385,
		79, 1, 0, 0, 0, 386, 388, 7, 0, 0, 0, 387, 386, 1, 0, 0, 0, 388, 389, 1,
		0, 0, 0, 389, 387, 1, 0, 0, 0, 389, 390, 1, 0, 0, 0, 390, 81, 1, 0, 0,
		0, 47, 85, 94, 104, 116, 131, 141, 144, 151, 159, 166, 170, 174, 179, 182,
		185, 190, 196, 204, 211, 217, 225, 232, 238, 241, 253, 272, 281, 292, 298,
		302, 305, 311, 316, 319, 322, 334, 337, 341, 348, 358, 362, 370, 377, 379,
		381, 384, 389,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// ApiParserParserInit initializes any static state used to implement ApiParserParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewApiParserParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ApiParserParserInit() {
	staticData := &apiparserParserStaticData
	staticData.once.Do(apiparserParserInit)
}

// NewApiParserParser produces a new parser instance for the optional input antlr.TokenStream.
func NewApiParserParser(input antlr.TokenStream) *ApiParserParser {
	ApiParserParserInit()
	this := new(ApiParserParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &apiparserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "ApiParser.g4"

	return this
}

// ApiParserParser tokens.
const (
	ApiParserParserEOF           = antlr.TokenEOF
	ApiParserParserT__0          = 1
	ApiParserParserT__1          = 2
	ApiParserParserT__2          = 3
	ApiParserParserT__3          = 4
	ApiParserParserT__4          = 5
	ApiParserParserT__5          = 6
	ApiParserParserT__6          = 7
	ApiParserParserT__7          = 8
	ApiParserParserT__8          = 9
	ApiParserParserT__9          = 10
	ApiParserParserT__10         = 11
	ApiParserParserT__11         = 12
	ApiParserParserT__12         = 13
	ApiParserParserATDOC         = 14
	ApiParserParserATHANDLER     = 15
	ApiParserParserATDESC        = 16
	ApiParserParserINTERFACE     = 17
	ApiParserParserATSERVER      = 18
	ApiParserParserATTYPE        = 19
	ApiParserParserAS            = 20
	ApiParserParserWS            = 21
	ApiParserParserCOMMENT       = 22
	ApiParserParserLINE_COMMENT  = 23
	ApiParserParserSTRING        = 24
	ApiParserParserRAW_STRING    = 25
	ApiParserParserLINE_VALUE    = 26
	ApiParserParserID            = 27
	ApiParserParserLetterOrDigit = 28
)

// ApiParserParser rules.
const (
	ApiParserParserRULE_api              = 0
	ApiParserParserRULE_spec             = 1
	ApiParserParserRULE_syntaxLit        = 2
	ApiParserParserRULE_importSpec       = 3
	ApiParserParserRULE_importLit        = 4
	ApiParserParserRULE_importBlock      = 5
	ApiParserParserRULE_importBlockValue = 6
	ApiParserParserRULE_importValue      = 7
	ApiParserParserRULE_infoSpec         = 8
	ApiParserParserRULE_svcommLit        = 9
	ApiParserParserRULE_typeSpec         = 10
	ApiParserParserRULE_typeLit          = 11
	ApiParserParserRULE_typeBlock        = 12
	ApiParserParserRULE_typeLitBody      = 13
	ApiParserParserRULE_typeBlockBody    = 14
	ApiParserParserRULE_atType           = 15
	ApiParserParserRULE_typeStruct       = 16
	ApiParserParserRULE_typeAlias        = 17
	ApiParserParserRULE_typeBlockStruct  = 18
	ApiParserParserRULE_typeBlockAlias   = 19
	ApiParserParserRULE_field            = 20
	ApiParserParserRULE_normalField      = 21
	ApiParserParserRULE_anonymousFiled   = 22
	ApiParserParserRULE_dataType         = 23
	ApiParserParserRULE_pointerType      = 24
	ApiParserParserRULE_mapType          = 25
	ApiParserParserRULE_arrayType        = 26
	ApiParserParserRULE_serviceSpec      = 27
	ApiParserParserRULE_atServer         = 28
	ApiParserParserRULE_serviceApi       = 29
	ApiParserParserRULE_serviceRoute     = 30
	ApiParserParserRULE_atDoc            = 31
	ApiParserParserRULE_atHandler        = 32
	ApiParserParserRULE_atDesc           = 33
	ApiParserParserRULE_route            = 34
	ApiParserParserRULE_body             = 35
	ApiParserParserRULE_replybody        = 36
	ApiParserParserRULE_kvLit            = 37
	ApiParserParserRULE_serviceName      = 38
	ApiParserParserRULE_path             = 39
	ApiParserParserRULE_pathItem         = 40
)

// IApiContext is an interface to support dynamic dispatch.
type IApiContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllSpec() []ISpecContext
	Spec(i int) ISpecContext

	// IsApiContext differentiates from other interfaces.
	IsApiContext()
}

type ApiContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyApiContext() *ApiContext {
	var p = new(ApiContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ApiParserParserRULE_api
	return p
}

func (*ApiContext) IsApiContext() {}

func NewApiContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ApiContext {
	var p = new(ApiContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_api

	return p
}

func (s *ApiContext) GetParser() antlr.Parser { return s.parser }

func (s *ApiContext) AllSpec() []ISpecContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISpecContext); ok {
			len++
		}
	}

	tst := make([]ISpecContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISpecContext); ok {
			tst[i] = t.(ISpecContext)
			i++
		}
	}

	return tst
}

func (s *ApiContext) Spec(i int) ISpecContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISpecContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISpecContext)
}

func (s *ApiContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ApiContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ApiContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitApi(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParserParser) Api() (localctx IApiContext) {
	this := p
	_ = this

	localctx = NewApiContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ApiParserParserRULE_api)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(85)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&135004160) != 0 {
		{
			p.SetState(82)
			p.Spec()
		}

		p.SetState(87)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ISpecContext is an interface to support dynamic dispatch.
type ISpecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SyntaxLit() ISyntaxLitContext
	SvcommLit() ISvcommLitContext
	ImportSpec() IImportSpecContext
	InfoSpec() IInfoSpecContext
	TypeSpec() ITypeSpecContext
	ServiceSpec() IServiceSpecContext

	// IsSpecContext differentiates from other interfaces.
	IsSpecContext()
}

type SpecContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySpecContext() *SpecContext {
	var p = new(SpecContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ApiParserParserRULE_spec
	return p
}

func (*SpecContext) IsSpecContext() {}

func NewSpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SpecContext {
	var p = new(SpecContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_spec

	return p
}

func (s *SpecContext) GetParser() antlr.Parser { return s.parser }

func (s *SpecContext) SyntaxLit() ISyntaxLitContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISyntaxLitContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISyntaxLitContext)
}

func (s *SpecContext) SvcommLit() ISvcommLitContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISvcommLitContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISvcommLitContext)
}

func (s *SpecContext) ImportSpec() IImportSpecContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IImportSpecContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IImportSpecContext)
}

func (s *SpecContext) InfoSpec() IInfoSpecContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInfoSpecContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInfoSpecContext)
}

func (s *SpecContext) TypeSpec() ITypeSpecContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeSpecContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeSpecContext)
}

func (s *SpecContext) ServiceSpec() IServiceSpecContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IServiceSpecContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IServiceSpecContext)
}

func (s *SpecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SpecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SpecContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitSpec(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParserParser) Spec() (localctx ISpecContext) {
	this := p
	_ = this

	localctx = NewSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ApiParserParserRULE_spec)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(94)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(88)
			p.SyntaxLit()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(89)
			p.SvcommLit()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(90)
			p.ImportSpec()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(91)
			p.InfoSpec()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(92)
			p.TypeSpec()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(93)
			p.ServiceSpec()
		}

	}

	return localctx
}
