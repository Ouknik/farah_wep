// init formData
var formData = {};

var _instanceSweetalert = new sweetalert();



function changeStatus(value , orderid) {
 
    formData["status"]  = value;
    formData["orderid"]  = orderid;
    _instanceSweetalert.showBusy("Bitte warten ...");
  
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
   
    if (response.error == 1) {
        setTimeout(function(){
            _instanceSweetalert.init(response.message, true, "Ok", true);
            _instanceSweetalert.setType("error");
            _instanceSweetalert.show();
        }, 100);
    } else {

        setTimeout(
            function() {
                _instanceSweetalert.init(response.message, "");
                _instanceSweetalert.setType("success");
                _instanceSweetalert.show();
                setTimeout(function () {
                    _instanceSweetalert.close(1000);
                }, 100);
                // updateInvestorSection_1_4();
                scenarioSectionsForUpdateElement.innerHTML = response.parsedTemplate

            },
        500);

        location.reload();
    }
}



//Model Search 

// init MODALS
var modalInstance;
var _instanceSweetalert = new sweetalert();
document.addEventListener('DOMContentLoaded', function() {
    modalInstance = M.Modal.init(document.getElementById('modalSearchFormDate'), {});
});

document.addEventListener('DOMContentLoaded', function() {
  modalInstance = M.Modal.init(document.getElementById('modalSearchByDate'), {});

});

// Get the orders table and all of its rows
const ordersTable = document.getElementById('ordersTable');
const rows = ordersTable.getElementsByTagName('tr');

// Get the button to show the calendar and the calendar div

const date = document.getElementById('date-calendar');

const showFilter = document.getElementById('showFilter');



const datebegin = document.getElementById('date-begin-calendar');


const dateend = document.getElementById('date-end-calendar');




const statueButton = document.getElementById('statue-button');

//const statue = document.getElementById('statue-input');
//const icon_statue = document.getElementById('icon-statue');


const statue = document.querySelectorAll('#statue-input .icon-button');

 const icon_statue = document.getElementById('statue-input');

 const dateFilterValueElement = document.getElementById('dateFilter');





















function clearFilter() {

  datebegin.value = ''; 
  dateend.value = ''; 
  date.value="";
  showFilter.style.display = 'none';
  lodingrows ();
  
}




// Add an event listener to selected by date range
function SerchByDate() {

  showFilter.style.display = 'block';
  
  dateFilterValueElement.textContent = date.value ;

  var selectedDate ="";
  if ( date.value !=""){
  var dateParts =  date.value.split("-");
  var selectedDate = dateParts[2] + "." + dateParts[1] + "." + dateParts[0];
}
let dataFound = false; // Variable to track if any data is found
  for (let i = 0; i < rows.length; i++) {
    const dateCell = rows[i].getElementsByTagName('td')[2];
    if (dateCell) {
      const orderDate = dateCell.textContent.trim();

      if(selectedDate==""){
        rows[i].style.display = '';
        dataFound = true;
      }
     else if (orderDate.slice(0, 10) === selectedDate) {
        rows[i].style.display = '';
        dataFound = true;
      } else {
       
        rows[i].style.display = 'none';
      }
    }
  }

// Display "No items found" if no data is found
if (!dataFound) {
  // Check if "No items found" element already exists
  const noItemsRow = document.getElementById('noItemsRow');

  if (!noItemsRow) {
    // Create the "No Order found" message elements
    const noItemsRow = document.createElement('tr');
    noItemsRow.id = 'noItemsRow';

    const noItemsCell = document.createElement('td');
    noItemsCell.textContent = 'No Order found';
    noItemsCell.colSpan = '6'; // Set the colspan to span across all columns in the table
    noItemsCell.style.textAlign = 'center'; // Center align the message

    noItemsRow.appendChild(noItemsCell);

    // Insert the "No items found" message at the end of the table
    const tbody = ordersTable.getElementsByTagName('tbody')[0];
    tbody.appendChild(noItemsRow);
  }
} else {
  // Remove the "No items found" message if it exists
  const noItemsRow = document.getElementById('noItemsRow');
  if (noItemsRow) {
    noItemsRow.remove();
  }
}



}



function SerchByTwoDate() {
  showFilter.style.display = 'block';
  
  dateFilterValueElement.textContent ="From "+ datebegin.value + " To "+ dateend.value;

  






  var dateParts =  datebegin.value.split("-");
  const beginDate =dateParts[2] + "." + dateParts[1] + "." + dateParts[0];

  var dateParts2 =dateend.value.split("-");
  const endDate =dateParts2[2] + "." + dateParts2[1] + "." + dateParts2[0];

 console.log(datebegin.value);
 console.log(beginDate);
 

  let dataFound = false; // Variable to track if any data is found

  for (let i = 0; i < rows.length; i++) {
    const dateCell = rows[i].getElementsByTagName('td')[2];

    if (dateCell) {
      const orderDate = dateCell.textContent.trim();
      

      if (beginDate === '' && endDate === '') {
        rows[i].style.display = ''; // Show all rows if no dates are selected
        dataFound = true;
      } else if (orderDate.slice(0, 10) >= beginDate && orderDate.slice(0, 10) <= endDate) {
        rows[i].style.display = ''; // Show rows that fall within the date range
        dataFound = true;
      } else {
        rows[i].style.display = 'none'; // Hide rows that are outside the date range
      }
    }
  }

  // Display "No items found" if no data is found
  if (!dataFound) {
    // Check if "No items found" element already exists
    const noItemsRow = document.getElementById('noItemsRow');

    if (!noItemsRow) {
      // Create the "No Order found" message elements
      const noItemsRow = document.createElement('tr');
      noItemsRow.id = 'noItemsRow';

      const noItemsCell = document.createElement('td');
      noItemsCell.textContent = 'No Order found';
      noItemsCell.colSpan = '6'; // Set the colspan to span across all columns in the table
      noItemsCell.style.textAlign = 'center'; // Center align the message

      noItemsRow.appendChild(noItemsCell);

      // Insert the "No items found" message at the end of the table
      const tbody = ordersTable.getElementsByTagName('tbody')[0];
      tbody.appendChild(noItemsRow);
    }
  } else {
    // Remove the "No items found" message if it exists
    const noItemsRow = document.getElementById('noItemsRow');
    if (noItemsRow) {
      noItemsRow.remove();
    }
  }
}



statue.forEach(button => {
  button.addEventListener('click', () => {
    const value = button.getAttribute('data-value');
   
    
    let dataFound = false; // Variable to track if any data is found
    
    for (let i = 0; i < rows.length; i++) {
      const statueCell = rows[i].getElementsByTagName('td')[5];
  
        
       


      if (statueCell) {
        const orderStatue = statueCell.textContent.trim();
        
       
        if (orderStatue === '') {
          rows[i].style.display = ''; // Show all rows if no dates are selected
          dataFound = true;
        } else if (orderStatue === value) {
          rows[i].style.display = ''; // Show rows that fall within the date range
          dataFound = true;
        } else {
          rows[i].style.display = 'none'; // Hide rows that are outside the date range
        }
      }
    }
  
    // Display "No items found" if no data is found
    if (!dataFound) {
      // Check if "No items found" element already exists
      const noItemsRow = document.getElementById('noItemsRow');
  
      if (!noItemsRow) {
        // Create the "No Order found" message elements
        const noItemsRow = document.createElement('tr');
        noItemsRow.id = 'noItemsRow';
  
        const noItemsCell = document.createElement('td');
        noItemsCell.textContent = 'No Order found';
        noItemsCell.colSpan = '6'; // Set the colspan to span across all columns in the table
        noItemsCell.style.textAlign = 'center'; // Center align the message
  
        noItemsRow.appendChild(noItemsCell);
  
        // Insert the "No items found" message at the end of the table
        const tbody = ordersTable.getElementsByTagName('tbody')[0];
        tbody.appendChild(noItemsRow);
      }
    } else {
      // Remove the "No items found" message if it exists
      const noItemsRow = document.getElementById('noItemsRow');
      if (noItemsRow) {
        noItemsRow.remove();
      }
    }


  });
});


function lodingrows (){
  for (let i = 0; i < rows.length; i++) {
    rows[i].style.display = '';
      
  }
}

