package cmd

import (
	"errors"
	"fmt"
	"image/jpeg"
	"os"

	gim "github.com/ozankasikci/go-image-merge"
	"github.com/spf13/cobra"
)

var (
	width  int
	height int
)

func init() {
	rootCmd.AddCommand(processCmd)
	processCmd.Flags().Int("x", width, "The width of the grid formed when merging the images")
	processCmd.Flags().Int("y", height, "The hight of the grid formed when merging the images")
}

var processCmd = &cobra.Command{
	Use:   "process <IMAGE_FILE1> <IMAGE_FILE2> ...",
	Short: "Process the images into one image",
	Long:  "Takes all of the image files processed and uses the order with the x and y length to make one image",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("Requires at least one image file be passed in")
		}
		err := mergeImages(args, width, height)
		return err
	},
}

func mergeImages(imageFileNames []string, width int, height int) error {
	grids := []*gim.Grid{}
	for _, imageName := range imageFileNames {
		grids = append(grids, &gim.Grid{ImageFilePath: imageName})
	}

	fmt.Sprintf("Width : %d \n", width)
	fmt.Sprintf("Height : %d \n", height)
	rgba, err := gim.New(grids, width, height).Merge()
	if err != nil {
		return err
	}

	file, err := os.Create("out.jpg")
	if err != nil {
		return err
	}

	err = jpeg.Encode(file, rgba, &jpeg.Options{Quality: 100})
	if err != nil {
		return err
	}
	return nil
}
