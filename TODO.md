# TODO List

### Developing
- [x] Command handler
- [x] Filter generators
- [x] Modifier generators
- [x] Rectangle generator
- [x] Add pixel distortion support
- [ ] Add chromatic distortion support
- [x] Add noise support
- [ ] Develop text render system
- [ ] Add pixel selection support
- [x] Add color selection support
- [x] Add HSV color format support
- [ ] Replace masks with pixel streams (channels)
- [ ] Rework for-in-for functions to accept streams of pixels
- [x] Add color channel extraction
- [ ] Add fairy tail effect (https://www.imgonline.com.ua/eng/fairy-tale-picture-effect.php)
- [ ] Add charcoal drawing effect (https://www.imgonline.com.ua/eng/charcoal-drawing.php)
- [ ] Add color number limitation
- [ ] Add posterization (https://www.imgonline.com.ua/eng/posterization.php)
- [ ] Add monochrome conversion (https://www.imgonline.com.ua/eng/monochrome-picture.php)
- [ ] Add Bayer filter (https://www.imgonline.com.ua/eng/bayer-color-filter-effect.php)
- [ ] Add engraving (https://www.imgonline.com.ua/eng/engraved-photo-effect.php)
- [ ] Add 8-bit conversion (https://www.imgonline.com.ua/eng/8bit-picture.php)
- [ ] Add Bokeh effect (https://www.imgonline.com.ua/eng/bokeh-effect.php)
- [ ] Add Cartoon filter (https://www.imgonline.com.ua/eng/cartoon-picture.php)
- [ ] Add Comic book filter (dot patterns)
- [x] Write the command parser
- [x] Write the interpreter tree
- [x] Write a standardization function to convert script format to command format
- [ ] Design a data structure to optimize keyword recognition
- [ ] Encapsulate all keywords and related handles into a map (keyword, handle)

### Implementing
- [x] Implement keyword recognition
- [x] Implement functionalities call
- [x] Implement axial blur
- [x] Implement median filter
- [ ] Implement circular blur
- [ ] Implement radial blur
- [ ] Implement sepia modifier (https://www.imgonline.com.ua/eng/add-effect-sepia.php)
- [ ] Add frame support (for overlay)
- [ ] Add vintage frame
- [ ] Implement fill selection
- [ ] Implement anchor point for rotation
- [x] Merge Distort and Mirror into Transform
- [ ] Optimize all interpolations by using corner cases (alpha = 0 or alpha = 1) or remove interpolations
- [x] Added Swirl function generator
- [ ] Treat literal atoms as one word
- [ ] Execute script from file
- [ ] Implement let/define handle
- [ ] Implement control structures
- [ ] Implement keyword approximation

### Refactoring
- [x] Check for each function to return origin centered images
- [ ] Make all error code checks for Nikrom to use bitwise AND
- [ ] Replace the current error system with one that would be able to hold multiple errors and the throwing command
- [x] Make all rectangle and point operations to have sub-handles in their respective handle OR make overloaded handles (add point point and add rectangle point)

### Testing
- [ ] Test cubic interpolation
- [ ] Unit tests for Median
- [ ] Unit tests for Shift
- [ ] Unit tests for Noise
- [ ] Unit tests for HSV
- [ ] Unit tests for Create Image
- [ ] Unit tests for Select

### Debugging
- [ ] The black border for the Swirl effect

### Documentation
- [ ] Add comments to every handle
