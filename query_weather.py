#-*- coding: utf-8 -*-
import os
import urllib.parse
import urllib.request

import requests
import simplejson


GEOCODE_BASE_URL = 'http://maps.googleapis.com/maps/api/geocode/json'


def geocode(address, sensor, **geo_args):
    geo_args.update({
        'address': address,
        'sensor': sensor  
    })

    url = GEOCODE_BASE_URL + '?' + urllib.parse.urlencode(geo_args)
    result = requests.get(GEOCODE_BASE_URL, params=geo_args)
    data = result.json()

    return data['results'][0]['geometry']['location']


def get_weather_info(api_key, lat, lng):
    url = 'https://api.forecast.io/forecast/{api_key}/{lat},{lng}?units=si'
    response = requests.get(url.format(
        api_key=api_key,
        lat=lat,
        lng=lng
    ))

    return response.json()


def main():
    api_key = os.environ.get('FORECAST_API_KEY', None)
    if api_key is None:
        print("Please define FORECAST_API_KEY")

    geo = geocode('Berlin', 'false')

    info = get_weather_info(api_key, geo['lat'], geo['lng'])
    print(info['daily']['summary'])


if __name__ == '__main__':
    main()
