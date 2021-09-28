package cmd

import "log"

// CommonOpts should be set externally, they shared across all commands
type CommonOpts struct {
	Version string

	Logger *log.Logger
}

// SetCommon satisfies CommonOptionsCommander interface and sets common option fields
// The method called by main for each command
func (c *CommonOpts) SetCommon(commonOpts CommonOpts) {
	c.Version = commonOpts.Version
	c.Logger = commonOpts.Logger
}

// CommonOptionsCommander extends flags.Commander with SetCommon
// All commands should implement this interfaces
type CommonOptionsCommander interface {
	SetCommon(commonOpts CommonOpts)
	Execute(args []string) error
}