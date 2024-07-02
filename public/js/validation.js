window.addEventListener("load", () => {
    const form = document.getElementById("loginForm");
    const inputs = form.querySelectorAll("input");

    // Reset form dan alert saat halaman di-refresh
    if (performance.navigation.type === 1) {
        inputs.forEach((input) => {
            input.value = "";
            input.classList.remove("is-invalid");
        });

        const errorAlert = document.getElementById("errorAlert");
        if (errorAlert) {
            errorAlert.style.display = "none";
        }
    }

    inputs.forEach((input) => {
        input.addEventListener("input", () => {
            if (input.value.trim() !== "") {
                input.classList.remove("is-invalid");
                if (input.nextElementSibling) {
                    input.nextElementSibling.style.display = "none";
                }
            } else {
                input.classList.add("is-invalid");
                if (input.nextElementSibling) {
                    input.nextElementSibling.style.display = "block";
                }
            }
        });
    });

    form.addEventListener("submit", (e) => {
        let isValid = true;
        inputs.forEach((input) => {
            if (input.value.trim() === "") {
                input.classList.add("is-invalid");
                if (input.nextElementSibling) {
                    input.nextElementSibling.style.display = "block";
                }
                isValid = false;
            } else {
                input.classList.remove("is-invalid");
                if (input.nextElementSibling) {
                    input.nextElementSibling.style.display = "none";
                }
            }
        });
        if (!isValid) {
            e.preventDefault();
        }
    });
});
