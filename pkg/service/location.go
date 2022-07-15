package service

import (
	"context"
	"github.com/ekomobile/dadata/v2"
	"github.com/ekomobile/dadata/v2/api/suggest"
	"math"
	"meishi_golang/pkg/repository"
	"meishi_golang/senti"
	"strconv"
)

//LocationService структура сервиса локации
type LocationService struct {
	repo repository.Location
}

//MyLocationService принимает репозиторий для работы с базой
func MyLocationService(repo repository.Location) *LocationService {
	return &LocationService{repo: repo}
}

//GetGeoCode Получение геокода по адресу
func (s *LocationService) GetGeoCode(text string) ([]senti.GeoCode, error) {
	api := dadata.NewSuggestApi()
	fromBound := suggest.Bound{Value: "region"}
	toBound := suggest.Bound{Value: "house"}
	params := suggest.RequestParams{
		Query:         text,
		Count:         10,
		Locations:     nil,
		RestrictValue: false,
		FromBound:     &fromBound,
		ToBound:       &toBound,
	}
	result, err := api.Address(context.Background(), &params)
	var res []senti.GeoCode
	var one senti.GeoCode
	metro := s.repo.GetMetro() //это плохое и временное решение
	//метро есть в запросе если заплатить за тариф
	var dist float64
	for _, v := range result {
		dist = 5000.0
		one.FullName = v.Value
		one.UnrestrictedValue = v.UnrestrictedValue
		one.GeoLat, _ = strconv.ParseFloat(v.Data.GeoLat, 64)
		one.GeoLon, _ = strconv.ParseFloat(v.Data.GeoLon, 64)
		one.Guid = v.Data.FiasID
		one.City = v.Data.City
		one.Region = v.Data.RegionTypeFull
		one.Country = v.Data.Country
		one.PostalCode = v.Data.PostalCode
		one.Settlement = v.Data.Settlement
		one.Street = v.Data.StreetTypeFull
		one.House = v.Data.House
		one.Block = v.Data.Block
		if v.Data.City == "Москва" {
			for _, l := range metro {
				d := distance(one.GeoLat, one.GeoLon, l.GeoLat, l.GeoLon)
				if d <= dist {
					dist = d
					one.Metro = l
					one.Metro.Distance = dist
				}
			}
		}
		res = append(res, one)
	}
	if err != nil {
		return res, err
	}
	return res, nil
}
func distance(geoLat1 float64, geoLon1 float64, geoLat2 float64, geoLon2 float64) float64 {
	dLat := deg2rad(geoLat2 - geoLat1)
	dLon := deg2rad(geoLon2 - geoLon1)
	a := math.Sin(dLat/2)*math.Sin(dLat/2) + math.Cos(deg2rad(geoLat1))*math.Cos(deg2rad(geoLat2))*math.Sin(dLon/2)*math.Sin(dLon/2)
	c := 2 * math.Asin(math.Sqrt(a))
	var r float64
	r = 6363564.142709288 //радиус земли по широте москвы???
	return r * c
}
func deg2rad(deg float64) float64 {
	return deg * (math.Pi / 180.0)
}
