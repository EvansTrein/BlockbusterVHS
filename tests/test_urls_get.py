from unittest import TestCase
from flaskr.config import app


class TestURLsGET(TestCase):
    def setUp(self):
        self.app = app.test_client()

    def test_home(self):
        response = self.app.get('/')
        self.assertEqual(response.status_code, 200)

    def test_create_client(self):
        response = self.app.get('/create_client')
        self.assertEqual(response.status_code, 200)

    def test_all_clients(self):
        response = self.app.get('/all_clients')
        self.assertEqual(response.status_code, 200)

    def test_create_vhstape(self):
        response = self.app.get('/create_vhstape')
        self.assertEqual(response.status_code, 200)

    def test_all_vhstapes(self):
        response = self.app.get('/all_vhstapes')
        self.assertEqual(response.status_code, 200)

    def test_create_rental(self):
        response = self.app.get('/create_rental')
        self.assertEqual(response.status_code, 200)

    def test_all_rentals(self):
        response = self.app.get('/all_rentals')
        self.assertEqual(response.status_code, 200)

    def test_404(self):
        response = self.app.get('/non-existent-page')
        self.assertEqual(response.status_code, 404)
