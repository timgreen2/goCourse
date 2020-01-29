| VALUE TYPES | REFERENCE TYPES |
| ----------- | --------------- |
| int         | slices          |
| float       | maps            |
| string      | channels        |
| bool        | pointers        |
| structs     | functions       |

For value types, you need to use pointers to change the values inside a function.

For reference types, you don't need to use pointers. You can modify them directly.
