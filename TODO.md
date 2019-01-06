### TODO List

# Developing
- [ ] Add gradient support
- [x] Add anti-aliasing support (can be achieved through a 2x2 blur)
- [ ] Add layering support
- [ ] Add rotation support
- [x] Add mirroring support
- [x] Add color modifiers support

# Implementing
- [x] Extend rescaling to bicubic interpolation
- [x] Extend filter appliance to support strength parameter
- [x] Add border conditions to filter appliance
- [x] Add support for any sizes of filters
- [x] Add color adder
- [x] Replace the coloring functions with only one transformation with matrix

# Refactoring
- [x] Move each group functions into a package and folder on its one
- [x] Ensure correct importing of auxiliary functions
- [ ] Refactor scaling to dynamically use interpolations
- [ ] Reduce number of interpolations arguments through a structure
- [ ] Move mirroring auxiliaries in a special directory

# Testing
- [ ] Write unit tests for filter appliance
- [ ] Write unit tests for rescaling
- [ ] Test cubic interpolation
- [ ] Test linear gradient

# Debugging
- [ ] Solve black border for blur filter

# Documentation
- [ ] Write comments for every function
