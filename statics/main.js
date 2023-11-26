document.addEventListener('DOMContentLoaded', function () {
    // Fetch courses and populate the list
    fetch('/courses')
        .then(response => response.json())
        .then(courses => {
            const coursesList = document.getElementById('coursesList');
            courses.forEach(course => {
                const listItem = document.createElement('li');
                listItem.textContent = `${course.title} - ${course.description}`;
                coursesList.appendChild(listItem);
            });
        });

    // Fetch courses and populate the enrollment form
    fetch('/courses')
        .then(response => response.json())
        .then(courses => {
            const enrollmentForm = document.getElementById('enrollmentForm');
            const courseIdSelect = document.getElementById('courseId');
            courses.forEach(course => {
                const option = document.createElement('option');
                option.value = course.id;
                option.textContent = `${course.title} - ${course.description}`;
                courseIdSelect.appendChild(option);
            });
        });

    // Handle enrollment form submission
    const enrollmentForm = document.getElementById('enrollmentForm');
    enrollmentForm.addEventListener('submit', function (event) {
        event.preventDefault();
        const courseId = document.getElementById('courseId').value;

        // Simplified logic for demonstration
        fetch('/enrollments', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ courseId }),
        })
            .then(response => response.json())
            .then(result => {
                alert(result.message);
            });
    });
});
