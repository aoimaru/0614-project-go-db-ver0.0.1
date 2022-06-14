
import requests



def main():
    response = requests.post(
        '0.0.0.0:8080/city', 
        data={
            "Name": "Yatsushiro",
            "CountryCode": "JPN",
            "District": "kumamoto",
            "Population": "20000"
        }
    )
    print(response.status_code)    # HTTPのステータスコード取得
    print(response.text)    # レスポンスのHTMLを文字列で取得


if __name__ == "__main__":
    main()