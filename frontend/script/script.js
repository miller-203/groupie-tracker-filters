const minSlider_creation = document.getElementById("min-creation-date");
const maxSlider_creation = document.getElementById("max-creation-date");
const minOutput_creation = document.getElementById("min-creation-date-value");
const maxOutput_creation = document.getElementById("max-creation-date-value");

minSlider_creation.oninput = function () {
    if (parseInt(minSlider_creation.value) >= parseInt(maxSlider_creation.value)) {
        minSlider_creation.value = maxSlider_creation.value - 1;
    }
    minOutput_creation.innerHTML = minSlider_creation.value;
};

maxSlider_creation.oninput = function () {
    if (parseInt(maxSlider_creation.value) <= parseInt(minSlider_creation.value)) {
        maxSlider_creation.value = parseInt(minSlider_creation.value) + 1;
    }
    maxOutput_creation.innerHTML = maxSlider_creation.value;
};

const minSlider_album = document.getElementById("min-first-album-date");
const maxSlider_album = document.getElementById("max-first-album-date");
const minOutput_album = document.getElementById("min-first-album-date-value");
const maxOutput_album = document.getElementById("max-first-album-date-value");

minSlider_album.oninput = function () {
    if (parseInt(minSlider_album.value) >= parseInt(maxSlider_album.value)) {
        minSlider_album.value = maxSlider_album.value - 1;
    }
    minOutput_album.innerHTML = minSlider_album.value;
};

maxSlider_album.oninput = function () {
    if (parseInt(maxSlider_album.value) <= parseInt(minSlider_album.value)) {
        maxSlider_album.value = parseInt(minSlider_album.value) + 1;
    }
    maxOutput_album.innerHTML = maxSlider_album.value;
};

const searchInput = document.getElementById("search");
const artistBoxes = document.querySelectorAll(".artists-cards");
const memberCheckboxes = document.querySelectorAll('input[name="members"]');
const filterButton = document.getElementById("filter-button");

function filterArtists() {
    const searchValue = searchInput.value.toLowerCase();
    const selectedMembers = Array.from(memberCheckboxes)
        .filter(checkbox => checkbox.checked)
        .map(checkbox => checkbox.value.toLowerCase());
    const minCreationDate = parseInt(minSlider_creation.value);
    const maxCreationDate = parseInt(maxSlider_creation.value);
    const minFirstAlbumDate = parseInt(minSlider_album.value);
    const maxFirstAlbumDate = parseInt(maxSlider_album.value);

    artistBoxes.forEach((card) => {
        const locations = card.getAttribute("data-location").toLowerCase();
        const members = card.getAttribute("data-members");

        const creationDate = parseInt(card.getAttribute("data-creation-date"));
        const firstAlbumDate = parseInt((card.getAttribute("data-first-album").split("-"))[2]);

        const matchesLocation = locations.includes(searchValue);
        const matchesMembers = selectedMembers.includes(members) || selectedMembers.length === 0

        const matchesCreationDate = creationDate >= minCreationDate && creationDate <= maxCreationDate;
        const matchesFirstAlbumDate = firstAlbumDate >= minFirstAlbumDate && firstAlbumDate <= maxFirstAlbumDate;

        if (matchesLocation && matchesMembers && matchesCreationDate && matchesFirstAlbumDate) {
            card.style.display = "";
        } else {
            card.style.display = "none";
        }
    });
}
searchInput.addEventListener("input", filterArtists);
filterButton.addEventListener("click", filterArtists);
