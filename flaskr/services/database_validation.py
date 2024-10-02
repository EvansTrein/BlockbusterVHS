from flaskr.db_models import VhsTape, Client, Rental
from re import fullmatch


def validateCreateVhs(data_in: dict) -> dict:
    answer = {"error": False, "error_text": ""}
    title = data_in["title"]
    year = data_in["year"]
    age_rating = data_in["age_rating"]
    count = data_in["count"]
    key_in = data_in["key_in"]

    if VhsTape.query.filter_by(title=title).first() and key_in == "create":
        answer["error"] = True
        answer["error_text"] = "There's already a movie like that"
        return answer
    elif not all((title, year, age_rating, count)):
        answer["error"] = True
        answer["error_text"] = "All fields must be filled in"
        return answer
    elif not 1900 <= int(data_in["year"]) <= 10000:
        answer["error"] = True
        answer["error_text"] = "Incorrect year format"
        return answer
    elif len(title) > 100 or len(age_rating) > 10:
        answer["error"] = True
        answer["error_text"] = (
            'The number of characters in the "name film" or "age rating" field is exceeded'
        )
        return answer
    else:
        return answer


def validateUpdateVhs(data_in: dict) -> str:
    error_text = ""
    obj_bd = data_in["obj_bd"]
    count = data_in["count"]

    is_vaild = validateCreateVhs(data_in)

    if not count.isdigit():
        error_text = "count can only be an integer"
        return error_text
    if int(count) < obj_bd.issued_to_clients:
        error_text = (
            "The total quantity may not be less than the sum of available and issued"
        )
        return error_text
    elif is_vaild["error"]:
        return is_vaild["error_text"]
    else:
        return error_text


def validateCreateClient(data_in: dict) -> dict:
    answer = {"error": False, "error_text": ""}
    name = data_in["name"]
    age = data_in["age"]
    phone = data_in["phone"]
    key_in = data_in["key_in"]

    if Client.query.filter_by(name=name).first() and key_in == "create":
        answer["error"] = True
        answer["error_text"] = "There's already a client like that"
        return answer
    elif not all((name, age, phone)):
        answer["error"] = True
        answer["error_text"] = "All fields must be filled in"
        return answer
    elif int(age) < 14:
        answer["error"] = True
        answer["error_text"] = "The client cannot be under 14 years of age"
        return answer
    elif not fullmatch(r"\+\d{10,20}", phone):
        answer["error"] = True
        answer["error_text"] = (
            "incorrect phone number (format: +country code... from 10 to 20 digits)"
        )
        return answer
    else:
        return answer


def validateUpdateClient(data_in: dict) -> str:
    error_text = ""

    is_vaild = validateCreateClient(data_in)

    if is_vaild["error"]:
        return is_vaild["error_text"]
    else:
        return error_text


def validateCreateRental(data_in: dict) -> dict:
    answer = {"error": False, "error_text": ""}

    client_id = data_in["client_id"]
    vhs_tape_id = data_in["vhs_tape_id"]
    client = Client.query.get(client_id)
    vhs = VhsTape.query.get(vhs_tape_id)
    existing_rental = Rental.query.filter_by(
        title_vhs=vhs.title, client_name=client.name
    ).first()

    if (client is None) and (vhs is None):
        answer["error"] = True
        answer["error_text"] = "The movie and client with those ID's do not exist"
        return answer
    elif not client:
        answer["error"] = True
        answer["error_text"] = "Client with this ID does not exist"
        return answer
    elif not vhs:
        answer["error"] = True
        answer["error_text"] = "A movie with that ID does not exist"
        return answer
    elif existing_rental:
        answer["error"] = True
        answer["error_text"] = "This customer has already rented this movie"
        return answer
    elif vhs.available_quantity == 0:
        answer["error"] = True
        answer["error_text"] = "Out of stock"
        return answer
    else:
        return answer
