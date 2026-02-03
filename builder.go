package gophitor

import (
	"time"
)

// gophitor.NewEntry("user.login", gophitor.CategoryAuthentication).
// Tenat("sdasdasd-uuid").
// Actor("uuuid", "user").
// Resource("session", "sess-abc").         // → ResourceBuilder
// Success().                               // → ReadyBuilder
// Log(ctx, logger)

type EntryBuilder struct {
	entry *Entry
}

type EntryBuilderWithTenant struct {
	entry *Entry
}

// TODO: This should not be a stand alone function but instead be part of the base struct
func NewEntry() *EntryBuilderWithTenant {
	return &EntryBuilderWithTenant{
		entry: &Entry{
			Timestamp: time.Now().UTC(),
			Metadata:  make(map[string]any),
		},
	}
}

func NewEntryWithTime(timestamp time.Time) *EntryBuilderWithTenant {
	return &EntryBuilderWithTenant{
		entry: &Entry{
			Timestamp: timestamp.UTC(),
			Metadata:  make(map[string]any),
		},
	}
}

func (b *EntryBuilderWithTenant) Workspace(workspaceID string) *EntryBuilderWithTenant {
	b.entry.WorkspaceID = workspaceID
	return &EntryBuilderWithTenant{entry: b.entry}
}

func (b *EntryBuilderWithTenant) Tenat(tenantID string) *ActorBuilder {
	b.entry.TenantID = tenantID
	return &ActorBuilder{entry: b.entry}
}

type ActorBuilder struct {
	entry *Entry
}

func (b *ActorBuilder) Actor(id string, actorType ActorType) *ActionBuilder {
	b.entry.ActorID = id
	b.entry.ActorType = actorType
	return &ActionBuilder{entry: b.entry}
}

type ActionBuilder struct {
	entry *Entry
}

func (b ActionBuilder) Action(action ActionKind) ResourceBuilder {
	b.entry.ActionKind = action
	return ResourceBuilder{entry: b.entry}
}

type ResourceBuilder struct {
	entry *Entry
}

func (b *ResourceBuilder) Resource(resourceID, resourceType string) ActionStateBuilder {
	b.entry.ResourceID = resourceID
	b.entry.ResourceType = resourceType
	return ActionStateBuilder{b.entry}
}

func (b *ResourceBuilder) ResourceWithName(resourceID, resourceType, resourceName string) ActionStateBuilder {
	b.entry.ResourceID = resourceID
	b.entry.ResourceType = resourceType
	b.entry.ResouceName = resourceName
	return ActionStateBuilder{b.entry}
}

type ActionStateBuilder struct {
	entry *Entry
}

func (b *ActionStateBuilder) Sucess() ReadyStateBuilder {
	b.entry.ActionStatus = StatusSuccesded
	return ReadyStateBuilder{
		b.entry,
	}
}

func (b *ActionStateBuilder) Failed() ReadyStateBuilder {
	b.entry.ActionStatus = StatusFailed
	return ReadyStateBuilder{
		b.entry,
	}
}

func (b *ActionStateBuilder) Denaied() ReadyStateBuilder {
	b.entry.ActionStatus = StatusDenied
	return ReadyStateBuilder{
		b.entry,
	}
}

// This is the final struct that actual commits the state
type ReadyStateBuilder struct {
	entry *Entry
}
