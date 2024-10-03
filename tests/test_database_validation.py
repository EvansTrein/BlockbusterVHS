from unittest import TestCase
from flaskr.services.database_validation import validateCreateVhs, validateCreateClient


class TestValidDate(TestCase):

    def test_validateCreateVhs(self):
        data_list = ({"title": "", "year": "", "age_rating": "", "count": "", "key_in": ""},
                    {"title": "", "year": "test", "age_rating": "", "count": "test", "key_in": ""},
                    {"title": "tets", "year": "900", "age_rating": "tets", "count": "tets", "key_in": "tets"},
                    {"title": "tets", "year": "not_digits", "age_rating": "tets", "count": "tets", "key_in": "tets"},
                    {"title": "tets", "year": "2001", "age_rating": "X"*12, "count": "tets", "key_in": "tets"},
                    {"title": 'X'*101, "year": "2001", "age_rating": "tets", "count": "tets", "key_in": "tets"}
                    )

        expected_response_list = ({"error": True, "error_text": 'All fields must be filled in'},
                                {"error": True, "error_text": 'All fields must be filled in'},
                                {"error": True, "error_text": 'Incorrect year'},
                                {"error": True, "error_text": 'Incorrect year'},
                                {"error": True, "error_text": 'The number of characters in the "name film" or "age rating" field is exceeded'},
                                {"error": True, "error_text": 'The number of characters in the "name film" or "age rating" field is exceeded'}
                                )

        for data, resp in zip(data_list, expected_response_list):
            self.assertEqual(validateCreateVhs(data), resp)
        
    
    def test_validateCreateClient(self):
        data_list = ({'name': '', 'age': '', 'phone': '', 'key_in': ''},
                    {'name': '', 'age': '25', 'phone': '+79000000000', 'key_in': ''},
                    {'name': 'test', 'age': '10', 'phone': '+79000000000', 'key_in': ''},
                    {'name': 'test', 'age': '20', 'phone': '45464353343245', 'key_in': ''},
                    {'name': 'test', 'age': '20', 'phone': '+79000', 'key_in': ''},
                    {'name': 'test', 'age': '20', 'phone': '+7900000000000000000000', 'key_in': ''}
        )

        expected_response_list = ({"error": True, "error_text": 'All fields must be filled in'}, 
                                {"error": True, "error_text": 'All fields must be filled in'},
                                {"error": True, "error_text": 'The client cannot be under 14 years of age'},
                                {"error": True, "error_text": 'incorrect phone number (format: +country code... from 10 to 20 digits)'},
                                {"error": True, "error_text": 'incorrect phone number (format: +country code... from 10 to 20 digits)'},
                                {"error": True, "error_text": 'incorrect phone number (format: +country code... from 10 to 20 digits)'}
                                )
        
        for data, resp in zip(data_list, expected_response_list):
            self.assertEqual(validateCreateClient(data), resp)