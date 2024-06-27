let showHidePasswordTogglers = null;

try {
	showHidePasswordTogglers = document.querySelectorAll(".toggle-show-hide-password");
} catch (e) {console.log(e);}


if (showHidePasswordTogglers != null) {
	showHidePasswordTogglers.forEach(function(elem) {
	    elem.addEventListener("click", function(e) {
	        let passwordElement = this.parentNode;
	        let inputType = passwordElement.querySelector("input").type;
	        if (inputType == "text") {
	        	passwordElement.querySelector("input").type = "password";
	        	this.querySelector("i").textContent = "visibility";
	        } else {
	        	passwordElement.querySelector("input").type = "text";
	        	this.querySelector("i").textContent = "visibility_off";
	        }
	    });
	});
}



