document.getElementById("year").innerText = new Date().getFullYear();

const hamburger = document.querySelector("#toggle-btn");

hamburger.addEventListener("click", function () {
    document.querySelector("#sidebar").classList.toggle("expand");
});

// Button Scroll Back to Top
// Get the button
let mybutton = document.getElementById("myBtn");

// When the user scrolls down 20px from the top of the document, show the button
window.onscroll = function () {
    scrollFunction();
};

function scrollFunction() {
    if (document.body.scrollTop > 20 || document.documentElement.scrollTop > 20) {
        mybutton.style.display = "block";
    } else {
        mybutton.style.display = "none";
    }
}

// When the user clicks on the button, scroll to the top of the document smoothly
function topFunction() {
    // Check if smooth scrolling is supported
    if ("scrollBehavior" in document.documentElement.style) {
        window.scrollTo({ top: 0, behavior: "smooth" }); // Smooth scroll to top
    } else {
        window.scrollTo(0, 0); // Fallback for browsers that do not support smooth scrolling
    }
}

// Display the current year in the footer
document.getElementById("year").textContent = new Date().getFullYear();
