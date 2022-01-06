package feed

import (
	"context"
	"log"
	"project-management-demo-backend/ent"
)

// Color generates color data
func Color(ctx context.Context, client *ent.Client) {
	_, err := client.Color.Delete().Exec(ctx)
	if err != nil {
		log.Fatalf("failed to delete color: %v", err)
	}

	ts := []ent.CreateColorInput{
		{Name: "white", Color: "white", Hex: "#FFFFFF"},
		{Name: "black", Color: "black", Hex: "#000000"},
		{Name: "gray.50", Color: "gray.50", Hex: "#F7FAFC"},
		{Name: "gray.100", Color: "gray.100", Hex: "#EDF2F7"},
		{Name: "gray.200", Color: "gray.200", Hex: "#E2E8F0"},
		{Name: "gray.300", Color: "gray.300", Hex: "#CBD5E0"},
		{Name: "gray.400", Color: "gray.400", Hex: "#A0AEC0"},
		{Name: "gray.500", Color: "gray.500", Hex: "#718096"},
		{Name: "gray.600", Color: "gray.600", Hex: "#4A5568"},
		{Name: "gray.700", Color: "gray.700", Hex: "#2D3748"},
		{Name: "gray.800", Color: "gray.800", Hex: "#1A202C"},
		{Name: "gray.900", Color: "gray.900", Hex: "#171923"},

		{Name: "red.50", Color: "red.50", Hex: "#FFF5F5"},
		{Name: "red.100", Color: "red.100", Hex: "#FED7D7"},
		{Name: "red.200", Color: "red.200", Hex: "#FEB2B2"},
		{Name: "red.300", Color: "red.300", Hex: "#FC8181"},
		{Name: "red.400", Color: "red.400", Hex: "#F56565"},
		{Name: "red.500", Color: "red.500", Hex: "#E53E3E"},
		{Name: "red.600", Color: "red.600", Hex: "#C53030"},
		{Name: "red.700", Color: "red.700", Hex: "#9B2C2C"},
		{Name: "red.800", Color: "red.800", Hex: "#822727"},
		{Name: "red.900", Color: "red.900", Hex: "#63171B"},

		{Name: "orange.50", Color: "orange.50", Hex: "#FFFAF0"},
		{Name: "orange.100", Color: "orange.100", Hex: "#FEEBC8"},
		{Name: "orange.200", Color: "orange.200", Hex: "#FBD38D"},
		{Name: "orange.300", Color: "orange.300", Hex: "#F6AD55"},
		{Name: "orange.400", Color: "orange.400", Hex: "#ED8936"},
		{Name: "orange.500", Color: "orange.500", Hex: "#DD6B20"},
		{Name: "orange.600", Color: "orange.600", Hex: "#C05621"},
		{Name: "orange.700", Color: "orange.700", Hex: "#9C4221"},
		{Name: "orange.800", Color: "orange.800", Hex: "#7B341E"},
		{Name: "orange.900", Color: "orange.900", Hex: "#652B19"},

		{Name: "yellow.50", Color: "yellow.50", Hex: "#FFFFF0"},
		{Name: "yellow.100", Color: "yellow.100", Hex: "#FEFCBF"},
		{Name: "yellow.200", Color: "yellow.200", Hex: "#FAF089"},
		{Name: "yellow.300", Color: "yellow.300", Hex: "#F6E05E"},
		{Name: "yellow.400", Color: "yellow.400", Hex: "#ECC94B"},
		{Name: "yellow.500", Color: "yellow.500", Hex: "#D69E2E"},
		{Name: "yellow.600", Color: "yellow.600", Hex: "#B7791F"},
		{Name: "yellow.700", Color: "yellow.700", Hex: "#975A16"},
		{Name: "yellow.800", Color: "yellow.800", Hex: "#744210"},
		{Name: "yellow.900", Color: "yellow.900", Hex: "#5F370E"},

		{Name: "green.50", Color: "green.50", Hex: "#F0FFF4"},
		{Name: "green.100", Color: "green.100", Hex: "#C6F6D5"},
		{Name: "green.200", Color: "green.200", Hex: "#9AE6B4"},
		{Name: "green.300", Color: "green.300", Hex: "#68D391"},
		{Name: "green.400", Color: "green.400", Hex: "#48BB78"},
		{Name: "green.500", Color: "green.500", Hex: "#38A169"},
		{Name: "green.600", Color: "green.600", Hex: "#2F855A"},
		{Name: "green.700", Color: "green.700", Hex: "#276749"},
		{Name: "green.800", Color: "green.800", Hex: "#22543D"},
		{Name: "green.900", Color: "green.900", Hex: "#1C4532"},

		{Name: "teal.50", Color: "teal.50", Hex: "#E6FFFA"},
		{Name: "teal.100", Color: "teal.100", Hex: "#B2F5EA"},
		{Name: "teal.200", Color: "teal.200", Hex: "#81E6D9"},
		{Name: "teal.300", Color: "teal.300", Hex: "#4FD1C5"},
		{Name: "teal.400", Color: "teal.400", Hex: "#38B2AC"},
		{Name: "teal.500", Color: "teal.500", Hex: "#319795"},
		{Name: "teal.600", Color: "teal.600", Hex: "#2C7A7B"},
		{Name: "teal.700", Color: "teal.700", Hex: "#285E61"},
		{Name: "teal.800", Color: "teal.800", Hex: "#234E52"},
		{Name: "teal.900", Color: "teal.900", Hex: "#1D4044"},

		{Name: "blue.50", Color: "blue.50", Hex: "#EBF8FF"},
		{Name: "blue.100", Color: "blue.100", Hex: "#BEE3F8"},
		{Name: "blue.200", Color: "blue.200", Hex: "#90CDF4"},
		{Name: "blue.300", Color: "blue.300", Hex: "#63B3ED"},
		{Name: "blue.400", Color: "blue.400", Hex: "#4299E1"},
		{Name: "blue.500", Color: "blue.500", Hex: "#3182CE"},
		{Name: "blue.600", Color: "blue.600", Hex: "#2B6CB0"},
		{Name: "blue.700", Color: "blue.700", Hex: "#2C5282"},
		{Name: "blue.800", Color: "blue.800", Hex: "#2A4365"},
		{Name: "blue.900", Color: "blue.900", Hex: "#1A365D"},

		{Name: "cyan.50", Color: "cyan.50", Hex: "#EDFDFD"},
		{Name: "cyan.100", Color: "cyan.100", Hex: "#C4F1F9"},
		{Name: "cyan.200", Color: "cyan.200", Hex: "#9DECF9"},
		{Name: "cyan.300", Color: "cyan.300", Hex: "#76E4F7"},
		{Name: "cyan.400", Color: "cyan.400", Hex: "#0BC5EA"},
		{Name: "cyan.500", Color: "cyan.500", Hex: "#00B5D8"},
		{Name: "cyan.600", Color: "cyan.600", Hex: "#00A3C4"},
		{Name: "cyan.700", Color: "cyan.700", Hex: "#0987A0"},
		{Name: "cyan.800", Color: "cyan.800", Hex: "#086F83"},
		{Name: "cyan.900", Color: "cyan.900", Hex: "#065666"},

		{Name: "purple.50", Color: "purple.50", Hex: "#FAF5FF"},
		{Name: "purple.100", Color: "purple.100", Hex: "#E9D8FD"},
		{Name: "purple.200", Color: "purple.200", Hex: "#D6BCFA"},
		{Name: "purple.300", Color: "purple.300", Hex: "#B794F4"},
		{Name: "purple.400", Color: "purple.400", Hex: "#9F7AEA"},
		{Name: "purple.500", Color: "purple.500", Hex: "#805AD5"},
		{Name: "purple.600", Color: "purple.600", Hex: "#6B46C1"},
		{Name: "purple.700", Color: "purple.700", Hex: "#553C9A"},
		{Name: "purple.800", Color: "purple.800", Hex: "#44337A"},
		{Name: "purple.900", Color: "purple.900", Hex: "#322659"},

		{Name: "pink.50", Color: "pink.50", Hex: "#FFF5F7"},
		{Name: "pink.100", Color: "pink.100", Hex: "#FED7E2"},
		{Name: "pink.200", Color: "pink.200", Hex: "#FBB6CE"},
		{Name: "pink.300", Color: "pink.300", Hex: "#F687B3"},
		{Name: "pink.400", Color: "pink.400", Hex: "#ED64A6"},
		{Name: "pink.500", Color: "pink.500", Hex: "#D53F8C"},
		{Name: "pink.600", Color: "pink.600", Hex: "#B83280"},
		{Name: "pink.700", Color: "pink.700", Hex: "#97266D"},
		{Name: "pink.800", Color: "pink.800", Hex: "#702459"},
		{Name: "pink.900", Color: "pink.900", Hex: "#521B41"},
	}
	bulk := make([]*ent.ColorCreate, len(ts))
	for i, t := range ts {
		bulk[i] = client.Color.Create().SetInput(t)
	}
	if _, err = client.Color.CreateBulk(bulk...).Save(ctx); err != nil {
		log.Fatalf("failed to feed color: %v", err)
	}
}
