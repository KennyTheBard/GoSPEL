### TODO List

# Developing
- [ ] Add gradient support
- [ ] Add anti-aliasing support
- [ ] Add layering support
- [ ] Add rotation support
- [ ] Add mirroring support

# Implementing
- [x] Extend rescaling to bicubic interpolation
- [x] Extend filter appliance to support strength parameter
- [x] Add border conditions to filter appliance
- [x] Add support for any sizes of filters

# Refactoring
- [x] Move each group functions into a package and folder on its one
- [x] Ensure correct importing of auxiliary functions
- [ ] Refactor scaling to dynamically use interpolations
- [ ] Reduce number of interpolations arguments through a structure

# Testing
- [ ] Write unit tests for filter appliance
- [ ] Write unit tests for rescaling
- [ ] Test cubic interpolation

# Debugging
- [ ] Solve black border for blur filter
