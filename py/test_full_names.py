#!/usr/bin/env python3
import unittest
from full_names import get_full_name

class NamesTestCase(unittest.TestCase):
    """ Tests for names.py """
    def test_first_last(self):
        full_name = get_full_name('janis', 'joplin')
        self.assertEqual(full_name, 'Janis Joplin')

#unittest.main()
