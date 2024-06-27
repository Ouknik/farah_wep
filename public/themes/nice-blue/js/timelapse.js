


var startTime, presentTime, elapsed, timekeepingId, closed;

function start(begin) {
  startTime = new Date(begin);
  // console.log(startTime);
};

function getPresentTime() {
  presentTime = new Date();
  //var timeDiff = presentTime - startTime; //in ms
  // strip the ms
  elapsed = (presentTime - startTime)/1000;

  // get seconds 
  //var seconds = Math.round(timeDiff);
  //console.log(seconds + " seconds");
  //console.log(Math.round(seconds/60) + " mins");

  const diff = {};

  //diff.days    = Math.floor(elapsed / 86400);
  diff.hours   = Math.floor(elapsed / 3600 % 24);
  diff.minutes = Math.floor(elapsed / 60 % 60);
  //diff.seconds = Math.floor(elapsed % 60);

  if (diff.hours < 10) {
    diff.hours = "0"+diff.hours
  }
  if (diff.minutes < 10) {
    diff.minutes = "0"+diff.minutes
  }
  /*
  if (diff.seconds < 10) {
    diff.seconds = "0"+diff.seconds
  }
  */


  document.getElementById("timelapsed").innerHTML = `${diff.hours}:${diff.minutes}`;

}


function updateElapsedTime() {
  
  ajax(`/task/updatetimekeeping/${timekeepingId}`, "get", {}, updateTime);

}

function updateTime(data) {
  var response = JSON.parse(data.response);
  //console.log(response);
  document.getElementById("lastupdated-container").innerHTML = response["time"];
}


// Document load or document ready function
document.addEventListener('DOMContentLoaded', function() {
  closed = document.getElementsByName("timekeeping_closed")[0].value;
  if (closed == 0) {
    timekeepingId = document.getElementsByName("timekeeping_track_id")[0].value;
  	start(document.getElementsByName("timekeeping_begin")[0].value);
    setInterval(getPresentTime, 6000);
    setInterval(updateElapsedTime, (60000 * 15));
  }
});