import $ from 'jquery';

function Controller(token) {

  this.token = token;

  var actualKeys = {
    'ArrowUp':false,
    'ArrowDown':false,
    'ArrowLeft':false,
    'ArrowRight':false,
    'Space':false
  }

  var serverKeys;

  window.addEventListener('keyup', function(event) {
    if (event.code in actualKeys) {
      actualKeys[event.code] = false
    }
  }, false);

  window.addEventListener('keydown', function(event) {
    if (event.code in actualKeys) {
      actualKeys[event.code] = true
    }
  }, false);

  setInterval(update, 100);


  function update() {
    if (JSON.stringify(serverKeys) != JSON.stringify(actualKeys)) {

      serverKeys = JSON.parse(JSON.stringify(actualKeys));

      $.ajax({
        type: "POST",
        url: "api/movement",
        headers: { Token: token },
        data: JSON.stringify(serverKeys)
      });

      //console.log("updated to "+JSON.stringify(serverKeys));
    }
  }
}

export {Controller}
