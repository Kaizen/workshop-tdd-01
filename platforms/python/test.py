from unittest.mock import MagicMock
import unittest


class ExampleClass:
    """A simple example class"""

    someValue = "Default value"

    def __init__(self, some_value):
        self.someValue = some_value
        # empty constructor

    def foo(self):
        return self.someValue


class MyTestCase(unittest.TestCase):

    def test_something(self):
        thing = ExampleClass("Foo")
        self.assertEqual("Foo", thing.foo())

        thing.foo = MagicMock(return_value="Bar")
        # Call function
        self.assertEqual("Bar", thing.foo())

        # Assert that it has been called
        thing.foo.assert_called()


if __name__ == '__main__':
    unittest.main()
