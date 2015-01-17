package filesystem

type Drive struct {
	DriveLetter string
}

// Drive factory function
func NewDrive(DriveLetter string) Drive {
	return Drive{DriveLetter}
}

func (d *Drive) Restore() {

}
