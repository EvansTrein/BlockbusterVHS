from .config import app, db
from .routes.home_and_others import route_home_and_others
from .routes.vhs import route_vhs
from .routes.client import route_client
from .routes.rental import route_rental

__all__ = [
    "app",
    "db",
    "route_home_and_others",
    "route_vhs",
    "route_client",
    "route_rental",
]
