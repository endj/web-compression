/*
    type UserInfo struct {
    	FieldA string
    	FieldB string
    	FieldC string
    	FieldD string
    	FieldE string
    	FieldF string
    	FieldG int
    	FieldH int
    	Text   string
    }
    */
const body = document.querySelector("body");

function generateTable(data, maxSize = 100) {
  const table = document.createElement("table");
  table.setAttribute("border", "1");

  const headerRow = document.createElement("tr");
  const headers = [
    "FieldA",
    "FieldB",
    "FieldC",
    "FieldD",
    "FieldE",
    "FieldF",
    "FieldG",
    "FieldH",
    "Text",
  ];

  headers.forEach((headerText) => {
    const th = document.createElement("th");
    th.textContent = headerText;
    headerRow.appendChild(th);
  });
  table.appendChild(headerRow);

  const rowsToRender = Math.min(maxSize, data.length);
  console.log(rowsToRender);
  for (let i = 0; i < rowsToRender; i++) {
    const user = data[i];
    const row = document.createElement("tr");

    Object.values(user).forEach((value) => {
      const td = document.createElement("td");
      td.textContent = isNaN(value) ? value.slice(0, 100) : value;
      row.appendChild(td);
    });

    table.appendChild(row);
  }

  body.appendChild(table);
}

async function fetchData() {
  const api = document.URL.split("/").pop().split(".")[0];

  const url = "http://localhost:8080/" + api;
  console.log("Calling API ", url);
  const response = await fetch(url);
  const json = await response.json();
  generateTable(json);
}
fetchData();
