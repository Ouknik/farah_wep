// default date format
try {
    if (dateFormat == null) dateFormat = "dd/mm/yyyy";
}catch (e) {
    var dateFormat = "dd/mm/yyyy";
}


try {
	// get the top nav user menu selector
	var topnavDropdownSelector = document.querySelector('.top-nav-user-menu-trigger');
} catch(e) {}

try {
    // get the top nav user menu selector
    M.Dropdown.init(document.querySelectorAll('.dropdowner'), {});
} catch(e) {}

// needed to format input values in 2.000,55 format
const germanNumberFormater = new Intl.NumberFormat('de-DE', { style: 'decimal', maximumFractionDigits: 2 });

// Document load or document ready function
document.addEventListener('DOMContentLoaded', function() {
	// instantiate top nav user menu drop down
	if (topnavDropdownSelector) {
	    M.Dropdown.init(topnavDropdownSelector, {'closeOnClick': true, 'onOpenStart': toggleTopnavDropdownActiveState, 'onCloseEnd': toggleTopnavDropdownActiveState});
	}
    // Initialize choose country in footer
    M.Dropdown.init(document.querySelector('.footer-choose-country-trigger'), {'closeOnClick': true});

    M.Sidenav.init(document.querySelectorAll('.sidenav'), {});

    // instantiate tooltips on all elements with class tooltipped
    M.Tooltip.init(document.querySelectorAll('.tooltipped'), {});

    // instanstiate dropdown select
    M.FormSelect.init(document.querySelectorAll('select'), {});

    // instanstiate datepickers
    // M.Datepicker.init(document.querySelectorAll('.datepicker'), {"format":"dd/mm/yyyy"});
    M.Datepicker.init(document.querySelectorAll('.datepicker'), {"format":dateFormat});

    // instanstiate tabs in cards
    M.Tabs.init(document.querySelector('.tabs'), {});

    // instanstiate collapsable
    M.Collapsible.init(document.querySelectorAll('.collapsible'), {});

    // Display global Toast Messages
    try {
        if (typeof(globalToastMessage) != "undefined" && globalToastMessage != "") {
            M.toast({html: globalToastMessage});
        }
    } catch (e) {
        console.log(e);
    }

    formatNumberToGermanLocale();

});


var toggleTopnavDropdownActiveState = function() {
	// toggle some style classes
	topnavDropdownSelector.classList.toggle("gradient-45deg-indigo-purple");
	topnavDropdownSelector.classList.toggle("gradient-shadow");
	topnavDropdownSelector.classList.toggle("white-text");
	// unfocus
	topnavDropdownSelector.blur();
}


// formData must be an array
// may also be an associated array
function ajax(url, type, formData, responseCallback) {
    // Creating the XMLHttpRequest object
    var request = new XMLHttpRequest();

    // remove the troubling leading slash (/)
    url = removeLeadingSlash(url);

    // Instantiating the request object
    request.open(type, url);
    request.setRequestHeader("Content-type", "application/x-www-form-urlencoded");

    // for some strange reason new FormData()
    // method does not work. Therefore, I descided
    // to handle things the old school way
    const parsedFormDataAsUrlString = convertFormDataToUrlString(formData);

    // Sending the request to the server
    request.send(parsedFormDataAsUrlString);
    // Defining event listener for readystatechange event
    request.onreadystatechange = function() {
        // Check if the request is compete and was successful
        //if(this.readyState === 4 && this.status === 200) {
        //    return this
        //}
        if (this.readyState === 4) {
        	responseCallback(this);
    	}
    };

}


// remove the troubling leading slash (/)
function removeLeadingSlash(url) {
    if (url.charAt(url.length - 1) == "/") url = url.slice(0, (url.length - 1));
    return url;
}



function convertFormDataToUrlString(formData) {
    let urlString = "";
    for (key in formData) {
        // console.log(key, formData[key]);
        urlString += `${key}=${formData[key]}&`;
    }
    return urlString;

}




function formatToCommaSeperatedString(stringToFormat, forceComma) {
const options = {
        minimumFractionDigits: 2,
        maximumFractionDigits: 2
    };
    return (!forceComma && !isFloat(stringToFormat)) ? stringToFormat : stringToFormat.toLocaleString('de-DE', options);
}

function isFloat(n){
    return Number(n) === n && n % 1 !== 0;
}

function convertAndParseFloat(valueToParse) {
    return parseFloat(valueToParse.replaceAll('.', '').replaceAll(',', '.'));
}


function isValidNumber(valueToCheck) {
    if (valueToCheck == "") return false;
    if (valueToCheck.match(/[^,.0-9]/)) return false;
    return true;
}

function formatNumberToGermanLocale(stringToFormat) {
    // console.log(stringToFormat)
    let formatedString = germanNumberFormater.format(stringToFormat);
    const splittedString = formatedString.split(",");
    if (splittedString.length > 1 && splittedString[1].length == 1) return formatedString + "0";
    return formatedString;
    // const inputElements = document.querySelectorAll('input[type="text"]');

    // for (var i = inputElements.length - 1; i >= 0; i--) {
    //     // console.log(inputElements[i]);
    //     if (inputElements[i].value != "" && !isNaN(inputElements[i].value)) {
    //         inputElements[i].value = germanNumberFormater.format(inputElements[i].value);
    //     }
    // }
}
