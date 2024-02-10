import requests

BASE_URL = 'http://localhost:8080/products'

def request_api(method, endpoint='', data=None):
    url = f"{BASE_URL}/{endpoint}".rstrip('/')
    response = requests.request(method, url, json=data)
    if response.status_code in [200, 201]:
        print(f"Operation successful. Response: {response.json() if response.text else ''}")
    else:
        print(f"Operation failed. Status code: {response.status_code}, Message: {response.text}")

def add_product(id, name, description, category, price, stock):
    request_api('POST', data={"id": id, "name": name, "description": description, "category": category, "price": price, "stock": stock})

def view_products():
    request_api('GET')

def get_product_by_id(product_id):
    request_api('GET', endpoint=str(product_id))

def delete_product_by_id(product_id):
    request_api('DELETE', endpoint=str(product_id))