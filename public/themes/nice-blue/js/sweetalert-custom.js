function sweetalert () {
//	var modalType, modalTitle, modalText, closeable, buttonText, buttonCloseModalOnClick, timer;
	var modalConfigOptions;
	var modalTriggered;
}



/**
	type: warning, success, error, info
	title: alert title
	text: alert text
	closeable: true/false
	buttonText: button caption
	buttonCloseModalOnClick: shall the alert close when button is clicked?
**/
sweetalert.prototype.init = function(title, text, closeable, buttonText, buttonCloseModalOnClick) {

	this.modalTriggered = false;

	// create a new config option object
	this.modalConfigOptions = {};

	// set title
	this.setConfigParam("title", title);
	// set text
	this.setConfigParam("text", text);
	// set closeable
	this.setConfigParam("closeOnClickOutside", closeable);
	// set closeOnEsc
	this.setConfigParam("closeOnEsc", closeable);

	// set button
	this.setButton(buttonText, buttonCloseModalOnClick);

}


sweetalert.prototype.showProgressbar = function(title) {

	// create a new config option object
	this.modalConfigOptions = {};

	// set text
	this.setConfigParam("text", "0%");

	this.initProgressbarAndShowBusy(title);
	
	this.createProgressbar();
}


sweetalert.prototype.showBusy = async function(title) {
	
	// create a new config option object
	this.modalConfigOptions = {};

	await this.initProgressbarAndShowBusy(title);

	this.createShowBusy();

	return;

}


sweetalert.prototype.initProgressbarAndShowBusy = function(title) {

	this.modalTriggered = false;

	// set title
	this.setConfigParam("title", title);
	// set closeable
	this.setConfigParam("closeOnClickOutside", false);
	// set closeOnEsc
	this.setConfigParam("closeOnEsc", false);

	// escaper button
	this.setButton();

	// show
	this.show();

}


sweetalert.prototype.createProgressbar = function() {
	// create dom element
	let div = document.createElement('div');
	// set template
	div.innerHTML = "<div class='progress grey lignten-4' style='width:80%; margin-left:auto; margin-right:auto;'><div id='progress-bar' style='width: 0%' class='determinate green darken-1'></div></div>";

	var parent = document.querySelector(".swal-modal");
	// before element
	var child = document.querySelector(".swal-text");
	parent.insertBefore(div, child);
}



sweetalert.prototype.createShowBusy = function() {
	// create dom element
	let div = document.createElement('div');
	// set template
	div.innerHTML = "<div class='progress blue lighten-4' style='width:80%; margin-left:auto; margin-right:auto;'><div class='indeterminate blue darken-1'></div></div>";      

	var parent = document.querySelector(".swal-modal");
	// before element
	var child = document.querySelector(".swal-text");
	parent.insertBefore(div, child);
}


sweetalert.prototype.updateProgressbar = function(progress) {
	this.updateText(`${progress}%`);
	try {
		document.querySelector("#progress-bar").style.width = progress + "%";
	} catch (e) {}
}


sweetalert.prototype.setConfigParam = function(param, value) {
	if ("" !== value) {
		this.modalConfigOptions[param] = value;
	}
}

sweetalert.prototype.setButton = function(buttonText, buttonCloseModalOnClick) {
	if (typeof(buttonText) != "undefined") {
		this.modalConfigOptions["button"] = {text: buttonText, closeModal: buttonCloseModalOnClick};
	} else {
		this.modalConfigOptions["buttons"] = false;
	}
}

sweetalert.prototype.setType = function(type) {
	this.setConfigParam("icon", type);
}

sweetalert.prototype.updateText = function(text) {
	if (!this.modalTriggered) return;
	setTimeout(
		function() {
			try {
				document.querySelector(".swal-text").innerHTML = text;
			} catch (e) {}
		},100
	);
}

sweetalert.prototype.close = function(timer) {

	var delay;
	if ('number' == typeof(timer)) {
		delay = timer;
	} else {
		delay = 100;
	}

	setTimeout(
		function() {
			//console.log(typeof(timer));
			//console.log(delay);
			this.modalTriggered = false;
			swal.close();
		}, delay
	);
}

sweetalert.prototype.show = function() {
	this.modalTriggered = true;
	swal(this.modalConfigOptions);

}


