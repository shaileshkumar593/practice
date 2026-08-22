"""
We are building a program to manage shared medical equipment bookings
for a clinic. The clinic has multiple pieces of equipment, such as
imaging rooms and portable devices, and staff can reserve them for
specific time periods.

Definitions:
* An "equipment" is an object representing a reservable item. It has
  properties for the ID and name.
* A "reservation" is an object representing a staff booking for a
  specific piece of equipment over a time period. Times are represented
  as integers in minutes from the start of the day.
* A "RentalManager" is the class used to manage all equipment and
  reservations.

To begin with, we present you with two tasks:
1-1) Read through and understand the code below. Please take as much
     time as necessary, and feel free to run the code.
1-2) One of the tests is failing due to a bug in the code. Make the
     necessary changes to RentalManager to fix the bug.
"""

import unittest
from enum import Enum


class ReservationStatus(Enum):
    """
    ReservationStatus represents the current state of a reservation.
    ACTIVE is the default status when a reservation is created.
    CANCELLED is set when a reservation is cancelled.
    """
    ACTIVE = 1
    CANCELLED = 2


class Equipment:
    """Data about a single piece of reservable medical equipment."""
    def __init__(self, equipment_id: int, name: str):
        self.equipment_id = equipment_id
        self.name = name

    def __str__(self):
        return "Equipment ID: %d, Name: %s" % (self.equipment_id, self.name)


class Reservation:
    """Data about a single equipment reservation."""
    def __init__(self, reservation_id: int, staff_name: str,
                 equipment_id: int, start_time: int, end_time: int):
        self.reservation_id = reservation_id
        self.staff_name = staff_name
        self.equipment_id = equipment_id
        self.start_time = start_time    # minutes from start of day, e.g. 480 = 8:00 AM
        self.end_time = end_time        # minutes from start of day
        self.status = ReservationStatus.ACTIVE

    def get_duration(self) -> int:
        return self.end_time - self.start_time

    def __eq__(self, other):
        if not isinstance(other, self.__class__):
            return False
        return self.reservation_id == other.reservation_id

    def __str__(self):
        return (
            "Reservation ID: %d, Staff: %s, Equipment ID: %d, "
            "Start: %d, End: %d, Status: %s"
            % (self.reservation_id, self.staff_name, self.equipment_id,
               self.start_time, self.end_time, self.status)
        )


class RentalManager:
    """
    Manages equipment reservations for the clinic.
    Equipment is identified by equipment_id.
    Reservations are stored and can be queried or cancelled.
    Two reservations conflict if they overlap in time for the same equipment.
    """
    def __init__(self):
        self.equipment_list = []
        self.reservations = []

    def add_equipment(self, equipment: Equipment):
        """Adds a piece of equipment to the clinic inventory."""
        self.equipment_list.append(equipment)

    def make_reservation(self, reservation: Reservation) -> bool:
        """
        Makes a reservation if the equipment is available for the requested period.
        Returns True if successfully reserved, False if the equipment is unavailable.
        """
        if not self._is_available(reservation.equipment_id,
                                  reservation.start_time,
                                  reservation.end_time):
            return False
        self.reservations.append(reservation)
        return True

    def _is_available(self, equipment_id: int, start_time: int, end_time: int) -> bool:
        for res in self.reservations:
            if res.equipment_id == equipment_id:
                if start_time < res.end_time and end_time > res.start_time:
                    return False
        return True

    def _is_cancelled(self, reservation_id):
        for res in self.reservations:
            if res.reservation_id == reservation_id:
                res.status = ReservationStatus.CANCELLED

    def cancel_reservation(self, reservation_id: int) -> bool:
        """
        Cancels an active reservation by ID.
        Returns True if found and cancelled, False otherwise.
        """
        self._is_cancelled(reservation_id)
        
        for res in self.reservations:
            if res.reservation_id == reservation_id:
                res.status = ReservationStatus.CANCELLED
                return True
        return False

    def get_reservations_for_equipment(self, equipment_id: int) -> list:
        """Returns all active reservations for a given piece of equipment."""
        return [res for res in self.reservations
                if res.equipment_id == equipment_id
                and res.status == ReservationStatus.ACTIVE]


class TestSuite(unittest.TestCase):
    # These tests are not meant to be exhaustive, and primarily show usage.

    def test_make_reservation(self):
        manager = RentalManager()
        res = Reservation(101, "Alice Johnson", 1, 480, 600)
        self.assertTrue(manager.make_reservation(res))
        self.assertEqual(len(manager.reservations), 1)

    def test_conflict_detected(self):
        manager = RentalManager()
        res1 = Reservation(101, "Alice Johnson", 1, 480, 600)
        res2 = Reservation(102, "Bob Smith", 1, 540, 660)
        manager.make_reservation(res1)
        self.assertFalse(manager.make_reservation(res2))

    def test_different_equipment_no_conflict(self):
        manager = RentalManager()
        res1 = Reservation(101, "Alice Johnson", 1, 480, 600)
        res2 = Reservation(102, "Bob Smith", 2, 480, 600)
        self.assertTrue(manager.make_reservation(res1))
        self.assertTrue(manager.make_reservation(res2))

    def test_cancelled_reservation_frees_slot(self):
        manager = RentalManager()
        res1 = Reservation(101, "Alice Johnson", 1, 480, 600)
        manager.make_reservation(res1)
        manager.cancel_reservation(101)
        # After cancellation, the same slot should be bookable again
        res2 = Reservation(102, "Bob Smith", 1, 480, 600)
        self.assertTrue(manager.make_reservation(res2))

    def test_cancel_reservation_success(self):
        manager = RentalManager()
        res = Reservation(101, "Alice Johnson", 1, 480, 600)
        manager.make_reservation(res)

        # Cancellation returns True and status is updated to CANCELLED
        self.assertTrue(manager.cancel_reservation(101))
        self.assertEqual(manager.reservations[0].status, ReservationStatus.CANCELLED)

        # Cancelled reservation remains in the collection — not removed
        self.assertEqual(len(manager.reservations), 1)

        # Other fields are unchanged after cancellation
        self.assertEqual(manager.reservations[0].reservation_id, 101)
        self.assertEqual(manager.reservations[0].start_time, 480)
        self.assertEqual(manager.reservations[0].end_time, 600)
        self.assertEqual(manager.reservations[0].equipment_id, 1)
        self.assertEqual(manager.reservations[0].staff_name, "Alice Johnson")

        # Cancelling a non-existent reservation returns False
        self.assertFalse(manager.cancel_reservation(999))


if __name__ == "__main__":
    unittest.main()
