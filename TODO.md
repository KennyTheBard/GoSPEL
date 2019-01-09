# TODO List

### Developing
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
- [x] Add cropping support

### Implementing
- [x] Extend resizing to bicubic interpolation
- [x] Extend filter appliance to support strength parameter
- [x] Add border conditions to filter appliance
- [x] Add support for any sizes of filters
- [x] Add color adder
- [x] Replace the coloring functions with only one transformation with matrix
- [x] Extend color modification to support target area
- [x] Extend all target area dependent functions to use masks
- [x] Add oval gradients
- [ ] Add safety measures in gradient functions

### Refactoring
- [x] Move each group functions into a package and folder on its one
- [x] Ensure correct importing of auxiliary functions
- [x] Reduce number of interpolations arguments through a structure
- [x] Rename ret image as trg image in rescale.go

### Testing
- [ ] Write unit tests for filter appliance
- [x] Write unit tests for resizing
- [ ] Write unit tests for color modification
- [ ] Write unit tests for gradient
- [x] Write unit tests for mirroring
- [ ] Write unit tests for layer merging
- [ ] Write unit tests for rotation
- [ ] Write unit tests for shearing
- [x] Write unit tests for cropping
- [ ] Write unit tests for opacity
- [ ] Test cubic interpolation
- [x] Test linear gradient

### Debugging
- [x] Solve black border for blur filter
- [ ] Solve the distortions of rotation algorithm
- [ ] Solve the error with PNG encoding
- [x] Solve the linear gradient bug

### Documentation
- [x] Write comments for every function
