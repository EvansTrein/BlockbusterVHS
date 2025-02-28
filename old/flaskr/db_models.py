from flaskr.config import db
from datetime import date


class VhsTape(db.Model):
    """
    movie database
    movie poster is stored as text - image links
    """

    __bind_key__ = "rentals"
    __tablename__ = "vhs_tapes"
    id_num = db.Column(db.Integer, primary_key=True)
    title = db.Column(db.String)
    year = db.Column(db.Integer)
    age_rating = db.Column(db.String)
    count = db.Column(db.Integer)
    available_quantity = db.Column(db.Integer, default=0)
    issued_to_clients = db.Column(db.Integer, default=0)
    poster_image = db.Column(db.String(255))

    def __init__(self, *args, **kwargs):
        """Initializes a VhsTape with the given arguments and sets the available quantity to the count parameter"""
        super().__init__(*args, **kwargs)
        self.available_quantity = self.count


class Client(db.Model):
    """
    client database
    """

    __bind_key__ = "rentals"
    __tablename__ = "clients"
    id_client = db.Column(db.Integer, primary_key=True)
    name = db.Column(db.String)
    age = db.Column(db.Integer)
    phone = db.Column(db.String)

    def __repr__(self) -> str:
        return f"{self.name}"


class Rental(db.Model):
    """
    rental database
    one-to-many ratio
    1 customer - can rent several different movies
    1 customer - cannot rent the same movie 2 times or more
    """

    __bind_key__ = "rentals"
    __tablename__ = "rentals"
    id_link = db.Column(db.Integer, primary_key=True)
    client_id = db.Column(db.Integer, db.ForeignKey("clients.id_client"))
    vhs_tape_id = db.Column(db.Integer, db.ForeignKey("vhs_tapes.id_num"))
    client = db.relationship("Client", backref=db.backref("rentals", lazy="dynamic"))
    vhs_tapes = db.relationship(
        "VhsTape", backref=db.backref("rentals", lazy="dynamic")
    )
    client_name = db.Column(db.String)
    title_vhs = db.Column(db.String)
    date_created = db.Column(db.Date)

    def __init__(self, client_id, vhs_tape_id):
        """
        Initializes a Rental with the given client_id and vhs_tape_id.
        Queries the Client and VhsTape tables to get the client name and the VHS tape title, and sets the date_created to today's date.
        :param client_id: The id of the client that the VHS tape is being rented to
        :param vhs_tape_id: The id of the VHS tape that is being rented
        :type client_id: int
        :type vhs_tape_id: int
        """
        self.client_id = client_id
        self.vhs_tape_id = vhs_tape_id
        self.client_name = Client.query.get(client_id).name
        self.title_vhs = VhsTape.query.get(vhs_tape_id).title
        self.date_created = date.today()
