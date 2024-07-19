# Web Scraper

A simple web scraper written in Go that uses the `colly` library for web scraping. This tool allows you to extract data from web pages based on a provided configuration file and optionally handle pagination.

## Features

- Scrape data from web pages based on configurable selectors.
- Support for pagination to scrape multiple pages.
- Output data in JSON format to a specified file.

## Prerequisites

- Go 1.18 or later
- `colly` library (included in the `go.mod` file)

## Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/yourusername/your-repository.git
    ```

2. Navigate to the project directory:

    ```bash
    cd your-repository
    ```

3. Install dependencies:

    ```bash
    go mod tidy
    ```

## Usage

The scraper requires a configuration file in JSON format. Below is the structure of the configuration file:

```json
{
    "url": "https://example.com",
    "container-selector": ".item",
    "pagination-selector": ".pagination a.next",
    "selectors": {
        "title": "h1",
        "image": "img.main-image"
    },
    "allow-pagination": true,
    "output": "output.json"
}
