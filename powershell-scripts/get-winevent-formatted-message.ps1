
get-winevent -providername $args[1] | Select-Object Message | format-table -wrap