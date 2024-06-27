// init formData
var formData = {};

var _instanceSweetalert = new sweetalert();

function filterByDate() {

    // Get the date value
     var selectedDate = document.getElementById("myDate");
     
     formData["Datetime"] = selectedDate.value; 
     
     // Loop through each row in the table
  $("#ordersTable tbody tr").each(function() {
    // Get the value in the second column (order date)
    var orderDate = $(this).find("td:nth-child(2)").text();

    // If the order date matches the selected date, show the row, otherwise hide it
    if (orderDate === selectedDate) {
      $(this).show();
    } else {
      $(this).hide();
    }
  });
 
}

function chqngeStatus(select , orderid) {
    formData["status"]  = select.value;
    formData["orderid"]  = orderid;
    _instanceSweetalert.showBusy("Bitte warten ...");
   // ajax("/reusables/addBatch", "post",formData , handleCalculated_reusable_success);  
   setTimeout(
    function() {
        ajax("/restaurant/salesoverview/changestatus", "POST", formData, handleSuccessChangeStatus);
    },500
);
       
  }

  
  const scenarioSectionsForUpdateElement = document.getElementById("switch-blocked");
function handleSuccessChangeStatus(responseData) {
    const response = JSON.parse(responseData.response);
    _instanceSweetalert.close();
     console.log(response)
    if (response.error == 1) {
        setTimeout(function(){
            _instanceSweetalert.init("Der Vorgang wurde erfolgreich abgeschlossen!", response.errorMessage, true, "Ok", true);
            _instanceSweetalert.setType("error");
            _instanceSweetalert.show();
        }, 100);
    } else {

        setTimeout(
            function() {
                _instanceSweetalert.init("Erfolgreich blockiert!", "");
                _instanceSweetalert.setType("success");
                _instanceSweetalert.show();
                setTimeout(function () {
                    _instanceSweetalert.close(1000);
                }, 100);
                // updateInvestorSection_1_4();
                scenarioSectionsForUpdateElement.innerHTML = response.parsedTemplate

            },
        500);
    }
}




// Get the orders table and all of its rows
const ordersTable = document.getElementById('ordersTable');
const rows = ordersTable.getElementsByTagName('tr');

// Get the button to show the calendar and the calendar div
const showCalendarButton = document.getElementById('show-calendar');
const calendar = document.getElementById('calendar');

// Show the calendar when the button is clicked
showCalendarButton.addEventListener('click', () => {
  calendar.style.display = 'block';
});

// Initialize the calendar
const calendarElement = document.createElement('input');
calendarElement.type = 'date';

// Add an event listener to the calendar element to filter the table when a date is selected
calendarElement.addEventListener('change', (event) => {
  const selectedDate = event.target.value;
  for (let i = 0; i < rows.length; i++) {
    const dateCell = rows[i].getElementsByTagName('td')[1];
    if (dateCell) {
      const orderDate = dateCell.textContent.trim();

      console.log(orderDate.slice(0, 10))
      console.log(selectedDate)
      if (orderDate.slice(0, 10) === selectedDate) {
        rows[i].style.display = '';
      } else {
        rows[i].style.display = 'none';
      }
    }
  }
});

// Add the calendar to the calendar div
calendar.appendChild(calendarElement);