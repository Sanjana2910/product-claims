INSERT INTO products (id, mcr_id, title, brand, category_name, raw_json)
VALUES (
    'cfe6aa75-5da8-44f5-b587-56857841ad9f',
    'IND0000000031',
    'Maggi Nutri-licious Masala Veg Atta Noodles',
    'Maggi',
    'Pasta, Soup, Noodles & Vermicelli',
    '{
        "id": "cfe6aa75-5da8-44f5-b587-56857841ad9f",
        "mcrId": "IND0000000031",
        "title": "Maggi Nutri-licious Masala Veg Atta Noodles",
        "brand": "Maggi",
        "shortDescription": "Source of fibre and iron. Safe for children.",
        "aboutItems": ["Source of Fibre & Iron", "Made with 20 Spices & Herbs"],
        "complianceInfo": {
            "country_of_origin": "India",
            "fssai_no": "10012011000168"
        },
        "attributes": {
            "diet_type": "Vegetarian",
            "nutritional_details": "Protein 8.0g, Fiber 5.0g, Iron 3.70mg"
        }
    }'
)
ON CONFLICT (id) DO NOTHING;