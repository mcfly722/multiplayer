function Display() {

  document.onreadystatechange = function () {
    $(document.body).css('padding', '0');
    $(document.body).css('margin', '0');
  }

  setInterval(update, 1000/60);

  function update() {

  }


}

export {Display}
