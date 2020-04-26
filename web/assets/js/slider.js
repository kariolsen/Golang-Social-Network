function plusSlides(section_id, length, n) {
    showSlides(section_id, length, n);
}

function showSlides(section_id, length, n) {
    var slideBox = document.getElementById(section_id).getElementsByClassName("mySlidesBox");
    var slides = document.getElementById(section_id).getElementsByClassName("mySlides");
    var slideBox_width = slideBox[0].style.width;
    var slide_width = slides[0].style.width;
    var slideBox_margin_left = slideBox[0].style.marginLeft;
    if (slideBox_margin_left == "") { slideBox_margin_left = "0%" }

    slideBox_width = eval(slideBox_width.slice(0, -1));
    slide_width = eval(slide_width.slice(0, -1));
    slideBox_margin_left = eval(slideBox_margin_left.slice(0, -1));

    var slideBox_left_limit = (slideBox_width * -1) + slide_width;
    var slideBox_right_limit = 0

    if (n == "1") {
        if (slideBox_margin_left < slideBox_right_limit) {
            slideBox_margin_left = slideBox_margin_left + slide_width;
        }
    } else if (n == "-1") {
        if (slideBox_margin_left > slideBox_left_limit) {
            slideBox_margin_left = slideBox_margin_left - slide_width;
        }
    }
    slideBox[0].style.marginLeft = slideBox_margin_left + "%";
}