
// Document load or document ready function
document.addEventListener('DOMContentLoaded', function() {

	var elems = document.querySelectorAll('.carousel');
    var instances = M.Carousel.init(elems, {
	    fullWidth: true,
	    indicators: true
	});
	initGetQrCode("/userqrcode/hasqrinfo", "bismillah")
})

var handleGetQrCode = function(response) {
	console.log(response);
}
function initGetQrCode(apiUrl, containerId) {
	fetch(apiUrl)
	.then(response => {
		if (response.ok && response.headers.get('Content-Type') === 'image/png') {
			return response.blob();
		} else {
			return response.json();
		}
	})
	.then(data => {
		if (data) {
			handleGetQrCode(data);
		} else {
			const imageUrl = URL.createObjectURL(blob);
			const container = document.getElementById(containerId);
			container.innerHTML = `<img src="${imageUrl}" alt="PNG Image">`;
		}
	})
	.catch(error => {
		console.error(error);
	});
}
// showPngImage("http://localhost:2020/qrcode?number=234978239", "bismillah");




const sliderContainer = document.querySelector(".slider-container");
const slider = document.querySelector(".slider");
const slides = document.querySelectorAll(".slider-item");
const slideWidth = slides[0].clientWidth;

function switchSlide(slideId) {
  // Get the slides-container element and all the slide elements inside it
  var container = document.getElementById('slides-container');
  var slides = container.getElementsByTagName('div');

  // Loop through all the slide elements
  for (var i = 0; i < slides.length; i++) {
    var slide = slides[i];
    // If the current slide is the one we want to show, fade it in
    if (slide.id === slideId) {
      fadeIn(slide);
    }
    // Otherwise, fade it out
    else {
      fadeOut(slide);
    }
  }
}

function fadeIn(element) {
  // Set the initial opacity to 0
  element.style.opacity = 0;

  // Show the element
  element.style.display = "block";

  // Define the animation
  var fadeInEffect = setInterval(function () {
    // Increase the opacity by 0.1 until it reaches 1
    if (element.style.opacity < 1) {
      element.style.opacity = parseFloat(element.style.opacity) + 0.1;
    } else {
      // Stop the animation
      clearInterval(fadeInEffect);
    }
  }, 50);
}

function fadeOut(element) {
  // Define the animation
  var fadeOutEffect = setInterval(function () {
    // Decrease the opacity by 0.1 until it reaches 0
    if (element.style.opacity > 0) {
      element.style.opacity = parseFloat(element.style.opacity) - 0.1;
    } else {
      // Hide the element
      element.style.display = "none";
      // Stop the animation
      clearInterval(fadeOutEffect);
    }
  }, 50);
}
