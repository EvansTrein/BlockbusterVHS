from unittest import TestCase
from flaskr.config import app


class TestTemplates(TestCase):
    def setUp(self):
        self.app = app.test_client()

    def test_home_page(self):
        response = self.app.get('/')
        self.assertIn('<title>BlockbusterVHS</title>', response.data.decode())
        self.assertIn('<a href="/">Home</a>', response.data.decode())
        self.assertIn('<a href="/all_vhstapes">All VHS</a>', response.data.decode())
        self.assertIn('<a href="/create_vhstape">Add VHS</a>', response.data.decode())
        self.assertIn('<a href="/all_clients">All clients</a>', response.data.decode())
        self.assertIn('<a href="/create_client">Add client</a>', response.data.decode())
        self.assertIn('<a href="/all_rentals">All rentals</a>', response.data.decode())
        self.assertIn('<a href="/create_rental">Add rental</a>', response.data.decode())



