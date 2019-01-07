### TODO List

# Developing
- [x] Add gradient support
- [x] Add anti-aliasing support (can be achieved through a 2x2 blur)
- [x] Add layer merging support
- [x] Add shearing support
- [ ] Add rotation support
- [x] Add mirroring support
- [x] Add color modifiers support
- [ ] Procedure chaining system
- [ ] Command handler
- [x] Add easy image copy

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
- [x] Reduce number of interpolations arguments through a structure
- [x] Rename ret image as trg image in rescale.go

# Testing
- [ ] Write unit tests for filter appliance
- [ ] Write unit tests for rescaling
- [ ] Write unit tests for color modification
- [ ] Write unit tests for gradient
- [ ] Write unit tests for mirroring
- [ ] Write unit tests for layer merging
- [ ] Test cubic interpolation
- [x] Test linear gradient

# Debugging
- [ ] Solve black border for blur filter

# Documentation
- [x] Write comments for every function
