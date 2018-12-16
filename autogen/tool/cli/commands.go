// Code generated by goagen v1.4.0, DO NOT EDIT.
//
// API "feedpushr": CLI Commands
//
// Command:
// $ goagen
// --design=github.com/ncarlier/feedpushr/design
// --out=$(GOPATH)/src/github.com/ncarlier/feedpushr/autogen
// --version=v1.4.0

package cli

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/goadesign/goa"
	goaclient "github.com/goadesign/goa/client"
	uuid "github.com/goadesign/goa/uuid"
	"github.com/ncarlier/feedpushr/autogen/client"
	"github.com/spf13/cobra"
	"log"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

type (
	// CreateFeedCommand is the command line data structure for the create action of feed
	CreateFeedCommand struct {
		// Comma separated list of tags
		Tags string
		// Feed URL
		URL         string
		PrettyPrint bool
	}

	// DeleteFeedCommand is the command line data structure for the delete action of feed
	DeleteFeedCommand struct {
		// Feed ID
		ID          string
		PrettyPrint bool
	}

	// GetFeedCommand is the command line data structure for the get action of feed
	GetFeedCommand struct {
		// Feed ID
		ID          string
		PrettyPrint bool
	}

	// ListFeedCommand is the command line data structure for the list action of feed
	ListFeedCommand struct {
		// Fetch limit
		Limit int
		// Page to fetch
		Page        int
		PrettyPrint bool
	}

	// StartFeedCommand is the command line data structure for the start action of feed
	StartFeedCommand struct {
		ID          string
		PrettyPrint bool
	}

	// StopFeedCommand is the command line data structure for the stop action of feed
	StopFeedCommand struct {
		ID          string
		PrettyPrint bool
	}

	// UpdateFeedCommand is the command line data structure for the update action of feed
	UpdateFeedCommand struct {
		// Feed ID
		ID string
		// Comma separated list of tags
		Tags        string
		PrettyPrint bool
	}

	// ListFilterCommand is the command line data structure for the list action of filter
	ListFilterCommand struct {
		PrettyPrint bool
	}

	// GetHealthCommand is the command line data structure for the get action of health
	GetHealthCommand struct {
		PrettyPrint bool
	}

	// GetOpmlCommand is the command line data structure for the get action of opml
	GetOpmlCommand struct {
		PrettyPrint bool
	}

	// UploadOpmlCommand is the command line data structure for the upload action of opml
	UploadOpmlCommand struct {
		PrettyPrint bool
	}

	// GetOutputCommand is the command line data structure for the get action of output
	GetOutputCommand struct {
		PrettyPrint bool
	}

	// PubPshbCommand is the command line data structure for the pub action of pshb
	PubPshbCommand struct {
		PrettyPrint bool
	}

	// SubPshbCommand is the command line data structure for the sub action of pshb
	SubPshbCommand struct {
		// A hub-generated random string
		HubChallenge string
		// The hub-determined number of seconds that the subscription will stay active before expiring
		HubLeaseSeconds int
		// The literal string "subscribe" or "unsubscribe"
		HubMode string
		// The topic URL given in the corresponding subscription request
		HubTopic    string
		PrettyPrint bool
	}

	// GetSwaggerCommand is the command line data structure for the get action of swagger
	GetSwaggerCommand struct {
		PrettyPrint bool
	}

	// GetVarsCommand is the command line data structure for the get action of vars
	GetVarsCommand struct {
		PrettyPrint bool
	}
)

// RegisterCommands registers the resource action CLI commands.
func RegisterCommands(app *cobra.Command, c *client.Client) {
	var command, sub *cobra.Command
	command = &cobra.Command{
		Use:   "create",
		Short: `Create a new feed`,
	}
	tmp1 := new(CreateFeedCommand)
	sub = &cobra.Command{
		Use:   `feed ["/v1/feeds"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp1.Run(c, args) },
	}
	tmp1.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp1.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "delete",
		Short: `Delete a feed`,
	}
	tmp2 := new(DeleteFeedCommand)
	sub = &cobra.Command{
		Use:   `feed ["/v1/feeds/ID"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp2.Run(c, args) },
	}
	tmp2.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp2.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "get",
		Short: `get action`,
	}
	tmp3 := new(GetFeedCommand)
	sub = &cobra.Command{
		Use:   `feed ["/v1/feeds/ID"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp3.Run(c, args) },
	}
	tmp3.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp3.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp4 := new(GetHealthCommand)
	sub = &cobra.Command{
		Use:   `health ["/v1/healthz"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp4.Run(c, args) },
	}
	tmp4.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp4.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp5 := new(GetOpmlCommand)
	sub = &cobra.Command{
		Use:   `opml ["/v1/opml"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp5.Run(c, args) },
	}
	tmp5.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp5.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp6 := new(GetOutputCommand)
	sub = &cobra.Command{
		Use:   `output ["/v1/output"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp6.Run(c, args) },
	}
	tmp6.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp6.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp7 := new(GetSwaggerCommand)
	sub = &cobra.Command{
		Use:   `swagger ["/v1/swagger.json"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp7.Run(c, args) },
	}
	tmp7.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp7.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp8 := new(GetVarsCommand)
	sub = &cobra.Command{
		Use:   `vars ["/v1/vars"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp8.Run(c, args) },
	}
	tmp8.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp8.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "list",
		Short: `list action`,
	}
	tmp9 := new(ListFeedCommand)
	sub = &cobra.Command{
		Use:   `feed ["/v1/feeds"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp9.Run(c, args) },
	}
	tmp9.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp9.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp10 := new(ListFilterCommand)
	sub = &cobra.Command{
		Use:   `filter ["/v1/filters"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp10.Run(c, args) },
	}
	tmp10.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp10.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "pub",
		Short: `Publication endpoint for PSHB hubs`,
	}
	tmp11 := new(PubPshbCommand)
	sub = &cobra.Command{
		Use:   `pshb ["/v1/pshb"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp11.Run(c, args) },
	}
	tmp11.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp11.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "start",
		Short: `Start feed aggregation`,
	}
	tmp12 := new(StartFeedCommand)
	sub = &cobra.Command{
		Use:   `feed ["/v1/feeds/ID/start"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp12.Run(c, args) },
	}
	tmp12.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp12.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "stop",
		Short: `Stop feed aggregation`,
	}
	tmp13 := new(StopFeedCommand)
	sub = &cobra.Command{
		Use:   `feed ["/v1/feeds/ID/stop"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp13.Run(c, args) },
	}
	tmp13.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp13.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "sub",
		Short: `Callback to validate the (un)subscription to the topic of a Hub`,
	}
	tmp14 := new(SubPshbCommand)
	sub = &cobra.Command{
		Use:   `pshb ["/v1/pshb"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp14.Run(c, args) },
	}
	tmp14.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp14.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "update",
		Short: `Update a feed`,
	}
	tmp15 := new(UpdateFeedCommand)
	sub = &cobra.Command{
		Use:   `feed ["/v1/feeds/ID"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp15.Run(c, args) },
	}
	tmp15.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp15.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "upload",
		Short: `Upload an OPML file to create feeds`,
	}
	tmp16 := new(UploadOpmlCommand)
	sub = &cobra.Command{
		Use:   `opml ["/v1/opml"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp16.Run(c, args) },
	}
	tmp16.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp16.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
}

func intFlagVal(name string, parsed int) *int {
	if hasFlag(name) {
		return &parsed
	}
	return nil
}

func float64FlagVal(name string, parsed float64) *float64 {
	if hasFlag(name) {
		return &parsed
	}
	return nil
}

func boolFlagVal(name string, parsed bool) *bool {
	if hasFlag(name) {
		return &parsed
	}
	return nil
}

func stringFlagVal(name string, parsed string) *string {
	if hasFlag(name) {
		return &parsed
	}
	return nil
}

func hasFlag(name string) bool {
	for _, arg := range os.Args[1:] {
		if strings.HasPrefix(arg, "--"+name) {
			return true
		}
	}
	return false
}

func jsonVal(val string) (*interface{}, error) {
	var t interface{}
	err := json.Unmarshal([]byte(val), &t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func jsonArray(ins []string) ([]interface{}, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []interface{}
	for _, id := range ins {
		val, err := jsonVal(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, val)
	}
	return vals, nil
}

func timeVal(val string) (*time.Time, error) {
	t, err := time.Parse(time.RFC3339, val)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func timeArray(ins []string) ([]time.Time, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []time.Time
	for _, id := range ins {
		val, err := timeVal(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, *val)
	}
	return vals, nil
}

func uuidVal(val string) (*uuid.UUID, error) {
	t, err := uuid.FromString(val)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func uuidArray(ins []string) ([]uuid.UUID, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []uuid.UUID
	for _, id := range ins {
		val, err := uuidVal(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, *val)
	}
	return vals, nil
}

func float64Val(val string) (*float64, error) {
	t, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func float64Array(ins []string) ([]float64, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []float64
	for _, id := range ins {
		val, err := float64Val(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, *val)
	}
	return vals, nil
}

func boolVal(val string) (*bool, error) {
	t, err := strconv.ParseBool(val)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func boolArray(ins []string) ([]bool, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []bool
	for _, id := range ins {
		val, err := boolVal(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, *val)
	}
	return vals, nil
}

// Run makes the HTTP request corresponding to the CreateFeedCommand command.
func (cmd *CreateFeedCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/v1/feeds"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.CreateFeed(ctx, path, cmd.URL, stringFlagVal("tags", cmd.Tags))
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *CreateFeedCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var tags string
	cc.Flags().StringVar(&cmd.Tags, "tags", tags, `Comma separated list of tags`)
	var url_ string
	cc.Flags().StringVar(&cmd.URL, "url", url_, `Feed URL`)
}

// Run makes the HTTP request corresponding to the DeleteFeedCommand command.
func (cmd *DeleteFeedCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/v1/feeds/%v", url.QueryEscape(cmd.ID))
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.DeleteFeed(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *DeleteFeedCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var id string
	cc.Flags().StringVar(&cmd.ID, "id", id, `Feed ID`)
}

// Run makes the HTTP request corresponding to the GetFeedCommand command.
func (cmd *GetFeedCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/v1/feeds/%v", url.QueryEscape(cmd.ID))
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.GetFeed(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *GetFeedCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var id string
	cc.Flags().StringVar(&cmd.ID, "id", id, `Feed ID`)
}

// Run makes the HTTP request corresponding to the ListFeedCommand command.
func (cmd *ListFeedCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/v1/feeds"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.ListFeed(ctx, path, intFlagVal("limit", cmd.Limit), intFlagVal("page", cmd.Page))
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ListFeedCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().IntVar(&cmd.Limit, "limit", 10, `Fetch limit`)
	cc.Flags().IntVar(&cmd.Page, "page", 1, `Page to fetch`)
}

// Run makes the HTTP request corresponding to the StartFeedCommand command.
func (cmd *StartFeedCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/v1/feeds/%v/start", url.QueryEscape(cmd.ID))
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.StartFeed(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *StartFeedCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var id string
	cc.Flags().StringVar(&cmd.ID, "id", id, ``)
}

// Run makes the HTTP request corresponding to the StopFeedCommand command.
func (cmd *StopFeedCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/v1/feeds/%v/stop", url.QueryEscape(cmd.ID))
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.StopFeed(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *StopFeedCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var id string
	cc.Flags().StringVar(&cmd.ID, "id", id, ``)
}

// Run makes the HTTP request corresponding to the UpdateFeedCommand command.
func (cmd *UpdateFeedCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/v1/feeds/%v", url.QueryEscape(cmd.ID))
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.UpdateFeed(ctx, path, stringFlagVal("tags", cmd.Tags))
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *UpdateFeedCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var id string
	cc.Flags().StringVar(&cmd.ID, "id", id, `Feed ID`)
	var tags string
	cc.Flags().StringVar(&cmd.Tags, "tags", tags, `Comma separated list of tags`)
}

// Run makes the HTTP request corresponding to the ListFilterCommand command.
func (cmd *ListFilterCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/v1/filters"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.ListFilter(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ListFilterCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
}

// Run makes the HTTP request corresponding to the GetHealthCommand command.
func (cmd *GetHealthCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/v1/healthz"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.GetHealth(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *GetHealthCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
}

// Run makes the HTTP request corresponding to the GetOpmlCommand command.
func (cmd *GetOpmlCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/v1/opml"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.GetOpml(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *GetOpmlCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
}

// Run makes the HTTP request corresponding to the UploadOpmlCommand command.
func (cmd *UploadOpmlCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/v1/opml"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.UploadOpml(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *UploadOpmlCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
}

// Run makes the HTTP request corresponding to the GetOutputCommand command.
func (cmd *GetOutputCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/v1/output"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.GetOutput(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *GetOutputCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
}

// Run makes the HTTP request corresponding to the PubPshbCommand command.
func (cmd *PubPshbCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/v1/pshb"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.PubPshb(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *PubPshbCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
}

// Run makes the HTTP request corresponding to the SubPshbCommand command.
func (cmd *SubPshbCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/v1/pshb"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.SubPshb(ctx, path, cmd.HubChallenge, cmd.HubMode, cmd.HubTopic, intFlagVal("hub.lease_seconds", cmd.HubLeaseSeconds))
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *SubPshbCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var hubChallenge string
	cc.Flags().StringVar(&cmd.HubChallenge, "hub.challenge", hubChallenge, `A hub-generated random string`)
	var hubLeaseSeconds int
	cc.Flags().IntVar(&cmd.HubLeaseSeconds, "hub.lease_seconds", hubLeaseSeconds, `The hub-determined number of seconds that the subscription will stay active before expiring`)
	var hubMode string
	cc.Flags().StringVar(&cmd.HubMode, "hub.mode", hubMode, `The literal string "subscribe" or "unsubscribe"`)
	var hubTopic string
	cc.Flags().StringVar(&cmd.HubTopic, "hub.topic", hubTopic, `The topic URL given in the corresponding subscription request`)
}

// Run makes the HTTP request corresponding to the GetSwaggerCommand command.
func (cmd *GetSwaggerCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/v1/swagger.json"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.GetSwagger(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *GetSwaggerCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
}

// Run makes the HTTP request corresponding to the GetVarsCommand command.
func (cmd *GetVarsCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/v1/vars"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.GetVars(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *GetVarsCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
}
