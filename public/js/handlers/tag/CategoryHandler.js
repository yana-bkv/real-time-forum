import {api} from '../../api/tag.js';

export async function CategoryHandler() {
    const result = await api.getTags();

    const container = document.getElementById("category-list");
    container.innerHTML = ""; // очищаем старые категории, если нужно

    result.forEach(tag => {
        const label = document.createElement("label");
        label.classList.add("category-checkbox");

        const checkbox = document.createElement("input");
        checkbox.type = "checkbox";
        checkbox.name = "categories";
        checkbox.value = tag.id;

        label.appendChild(document.createTextNode(tag.text));
        label.appendChild(checkbox);


        container.appendChild(label);
    });
}

export function showCategoryForm() {
    document.getElementById('new-category-form').style.display = 'block';
}

export async function createCategory() {
    const name = document.getElementById('new-category-name').value;
    const data = {
        text: name,
    }
    const category = await api.createTag(data);

    // Добавим новую категорию в select
    const select = document.getElementById('category-select');
    const option = document.createElement('option');
    option.value = category.id;
    option.text = category.name;
    option.selected = true; // автоматически выбираем её
    select.appendChild(option);

    window.location.href = '/create-post';
}